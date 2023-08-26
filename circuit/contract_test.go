package main

import (
	"bytes"
	"circuit/contract"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
)

type ExportSolidityTestSuiteGroth16 struct {
	suite.Suite

	// backend
	backend *backends.SimulatedBackend

	// verifier contract
	verifierContract *contract.GlacierVerifier

	// groth16 gnark objects
	vk      groth16.VerifyingKey
	pk      groth16.ProvingKey
	circuit Circuit
	r1cs    constraint.ConstraintSystem
}

func TestRunExportSolidityTestSuiteGroth16(t *testing.T) {
	suite.Run(t, new(ExportSolidityTestSuiteGroth16))
}

func (t *ExportSolidityTestSuiteGroth16) SetupTest() {

	const gasLimit uint64 = 4712388

	// setup simulated backend
	key, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	t.NoError(err, "init keyed transactor")

	genesis := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: big.NewInt(1000000000000000000)}, // 1 Eth
	}
	t.backend = backends.NewSimulatedBackend(genesis, gasLimit)

	// deploy verifier contract
	_, _, v, err := contract.DeployGlacierVerifier(auth, t.backend)
	t.NoError(err, "deploy verifier contract failed")
	t.verifierContract = v
	t.backend.Commit()

	t.r1cs, err = frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &t.circuit)
	t.NoError(err, "compiling R1CS failed")

	// read proving and verifying keys
	t.pk = groth16.NewProvingKey(ecc.BN254)
	{
		f, _ := os.Open("./contract/zkglacier.g16.pk")
		_, err = t.pk.ReadFrom(f)
		f.Close()
		t.NoError(err, "reading proving key failed")
	}
	t.vk = groth16.NewVerifyingKey(ecc.BN254)
	{
		f, _ := os.Open("./contract/zkglacier.g16.vk")
		_, err = t.vk.ReadFrom(f)
		f.Close()
		t.NoError(err, "reading verifying key failed")
	}

}

func (t *ExportSolidityTestSuiteGroth16) TestVerifyProof() {
	// mock a tranfer
	var rnd fr.Element
	rnd.SetUint64(uint64(0))
	src := rand.NewSource(int64(0))
	r := rand.New(src)
	owner, err := eddsa.GenerateKey(r)
	if err != nil {
		panic(err)
	}

	blockpreimage := sha256Sum([]byte("/block_data"))
	preBlockhash := sha256Sum([]byte(""))
	ts := uint64(time.Now().Unix())
	txs := genTestTx(16)
	bc := NewBlockCommitment(blockpreimage, preBlockhash, 1, ts, owner.PublicKey, txs)

	// sign the transfer on the client side.
	_, err = bc.SignWitness(*owner, hFunc)
	if err != nil {
		panic(err)
	}
	// create a valid proof
	blockHash := bc.Hash(hFunc)
	assignment := bc.Witness

	// witness creation
	witness, err := frontend.NewWitness(&assignment, ecc.BN254.ScalarField())
	t.NoError(err, "witness creation failed")

	// prove
	proof, err := groth16.Prove(t.r1cs, t.pk, witness)
	t.NoError(err, "proving failed")

	// ensure gnark (Go) code verifies it
	publicWitness, _ := witness.Public()
	err = groth16.Verify(proof, t.vk, publicWitness)
	t.NoError(err, "verifying failed")

	// get proof bytes
	const fpSize = 4 * 8
	var buf bytes.Buffer
	proof.WriteRawTo(&buf)
	proofBytes := buf.Bytes()

	// solidity contract inputs
	var (
		a     [2]*big.Int
		b     [2][2]*big.Int
		c     [2]*big.Int
		input [1]*big.Int
	)

	// proof.Ar, proof.Bs, proof.Krs
	a[0] = new(big.Int).SetBytes(proofBytes[fpSize*0 : fpSize*1])
	a[1] = new(big.Int).SetBytes(proofBytes[fpSize*1 : fpSize*2])
	b[0][0] = new(big.Int).SetBytes(proofBytes[fpSize*2 : fpSize*3])
	b[0][1] = new(big.Int).SetBytes(proofBytes[fpSize*3 : fpSize*4])
	b[1][0] = new(big.Int).SetBytes(proofBytes[fpSize*4 : fpSize*5])
	b[1][1] = new(big.Int).SetBytes(proofBytes[fpSize*5 : fpSize*6])
	c[0] = new(big.Int).SetBytes(proofBytes[fpSize*6 : fpSize*7])
	c[1] = new(big.Int).SetBytes(proofBytes[fpSize*7 : fpSize*8])

	// public witness
	input[0] = new(big.Int).SetBytes(blockHash)

	fmt.Println("contract", t.verifierContract)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(input)

	// call the contract
	res, err := t.verifierContract.VerifyProof(nil, a, b, c, input)
	fmt.Println("correctVerifyProof", err, res)
	if t.NoError(err, "calling verifier on chain gave error") {
		t.True(res, "calling verifier on chain didn't succeed")
	}

	// (wrong) public witness
	input[0] = new(big.Int).SetBytes([]byte("wrongHash!"))

	// call the contract should fail
	res, err = t.verifierContract.VerifyProof(nil, a, b, c, input)
	fmt.Println("wrongVerifyProof", err, res)
	if t.NoError(err, "calling verifier on chain gave error") {
		t.False(res, "calling verifier on chain succeed, and shouldn't have")
	}

	// gen mock blocks
	blks := []BlockInfo{}
	bcs := genTestBlockCommitment(t.T(), 5)
	for _, bc := range bcs {
		// generate the witness of the transfer
		witnessFull, err := frontend.NewWitness(&bc.Witness, ecc.BN254.ScalarField())
		if err != nil {
			panic(err)
		}

		// generate the proof of the transfer
		proofData, err := groth16.Prove(t.r1cs, t.pk, witnessFull)
		if err != nil {
			panic(err)
		}
		witnessPublic, err := witnessFull.Public()
		if err != nil {
			panic(err)
		}
		// Self-check the proof on the client side. Also, this proof will be submitted to ZK-Contract on the settlement EVM Chain.
		err = groth16.Verify(proofData, t.vk, witnessPublic)
		if err != nil {
			panic(err)
		}

		var buf bytes.Buffer
		proofData.WriteRawTo(&buf)
		proofBytes := buf.Bytes()

		blk := BlockInfo{
			BlockNumber:  bc.BlockNumber,
			BlockHash:    hexutil.Encode(bc.Hash(hFunc)),
			PreBlockHash: hexutil.Encode(bc.PreBlockHash),
			Timestamp:    bc.Timestamp,
			Commitment:   hexutil.Encode(proofBytes),
			DaId:         fmt.Sprintf("ar://arid%d?sha256=blocksha256", bc.BlockNumber),
		}

		fmt.Println(len(blk.BlockHash), len(blk.PreBlockHash), len(blk.Commitment))
		fmt.Println(blk.BlockNumber, blk.BlockHash, blk.Commitment)

		blks = append(blks, blk)

		if bc.BlockNumber == 1 {
			ioutil.WriteFile("../test/proof.bin", proofBytes, 0777)
			ioutil.WriteFile("../test/public.bin", bc.Hash(hFunc), 0777)
			fmt.Println("save blockNumber 1 prove&public")
		}
	}

	data, _ := json.Marshal(blks)
	ioutil.WriteFile("../test/mock_blocks.json", data, 0777)
}
