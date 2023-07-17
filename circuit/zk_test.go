package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

type BlockInfo struct {
	BlockNumber  uint64 `json:"blockNumber"`
	BlockHash    string `json:"blockHash"`    // hex
	PreBlockHash string `json:"preblockHash"` // hex
	Timestamp    uint64 `json:"timestamp"`
	Commitment   string `json:"commitment"` // hex
	DaId         string `json:"daid"`
}

func genTestBlockCommitment(t *testing.T, size int) (bcs []BlockCommitment) {
	var rnd fr.Element
	rnd.SetUint64(uint64(0))
	src := rand.NewSource(int64(0))
	r := rand.New(src)
	// a test account
	validator, err := eddsa.GenerateKey(r)
	if err != nil {
		panic(err)
	}

	// zero block
	preBlockhash := []byte("00000000000000000000000000000000")
	for i := 1; i < size; i++ {
		blockpreimage := sha256Sum([]byte(fmt.Sprintf("[tx01, tx02, tx03]/testdata,%d", i)))
		ts := uint64(time.Now().Unix())
		bc := NewBlockCommitment(blockpreimage, preBlockhash, uint64(i), ts, validator.PublicKey)
		// sign the transfer on the client side.
		_, err = bc.SignWitness(*validator, hFunc)
		if err != nil {
			t.Fatal(err)
		}
		// self-check sign on the client side.
		ok, err := bc.Verify(hFunc)
		if err != nil {
			t.Fatal(err)
		}

		if !ok {
			t.Fail()
		}

		preBlockhash = bc.Hash(hFunc)

		bcs = append(bcs, bc)
	}
	return
}

func TestZKSnarkProof(t *testing.T) {
	// step0: compile circuit
	var circuit Circuit

	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}

	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		panic(err)
	}

	blks := []BlockInfo{}
	bcs := genTestBlockCommitment(t, 5)
	for _, bc := range bcs {
		// generate the witness of the transfer
		witnessFull, err := frontend.NewWitness(&bc.Witness, ecc.BN254.ScalarField())
		if err != nil {
			panic(err)
		}

		// generate the proof of the transfer
		proofData, err := groth16.Prove(r1cs, pk, witnessFull)
		if err != nil {
			panic(err)
		}
		witnessPublic, err := witnessFull.Public()
		if err != nil {
			panic(err)
		}
		// Self-check the proof on the client side. Also, this proof will be submitted to ZK-Contract on the settlement EVM Chain.
		err = groth16.Verify(proofData, vk, witnessPublic)
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
		blks = append(blks, blk)
	}

	// data, _ := json.Marshal(blks)
	// ioutil.WriteFile("../test/mock_blocks.json", data, 0777)
}
