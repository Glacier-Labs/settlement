package main

import (
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/accumulator/merkle"
	"github.com/consensys/gnark/std/algebra/twistededwards"
	"github.com/consensys/gnark/std/hash/mimc"
	sigeddsa "github.com/consensys/gnark/std/signature/eddsa"
)

const MaxLeafNum = 16 // depth = 5
const TreeDepth = 5

type Circuit struct {
	Block              BlockConstraints
	ValidatorSignature sigeddsa.Signature
	ValidatorPubKey    sigeddsa.PublicKey
	BlockHash          frontend.Variable `gnark:",public"`
	TxRootHash         frontend.Variable `gnark:",public"`
	TxProofs           [MaxLeafNum]merkle.MerkleProof
	Txs                [MaxLeafNum]TxConstraints
}

func (c *Circuit) Define(api frontend.API) error {
	hFunc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	blockHash, err := verify(api, c.Block, hFunc)
	if err != nil {
		return err
	}

	api.AssertIsEqual(c.BlockHash, blockHash)

	curve, err := twistededwards.NewEdCurve(api, tedwards.BN254)
	if err != nil {
		return err
	}

	err = sigeddsa.Verify(curve, c.ValidatorSignature, blockHash, c.ValidatorPubKey, &hFunc)
	if err != nil {
		return err
	}

	err = verifyTxs(api, c.TxRootHash, c.TxProofs, c.Txs, hFunc)
	if err != nil {
		return err
	}

	return err
}

type BlockConstraints struct {
	BlockNumber   frontend.Variable
	BlockPreImage frontend.Variable
	PreBlockHash  frontend.Variable
	Timestamp     frontend.Variable
}

type TxConstraints struct {
	TxHash    frontend.Variable `gnark:",public"`
	Timestamp frontend.Variable
	Nonce     frontend.Variable
	Sgin      sigeddsa.Signature
	PubKey    sigeddsa.PublicKey
}

func verify(api frontend.API, t BlockConstraints, hFunc mimc.MiMC) (frontend.Variable, error) {
	hFunc.Reset()
	hFunc.Write(t.BlockNumber, t.BlockPreImage, t.PreBlockHash, t.Timestamp)
	hash := hFunc.Sum()
	return hash, nil
}

func verifyTxs(api frontend.API, TxRootHash frontend.Variable, proofs [MaxLeafNum]merkle.MerkleProof, txs [MaxLeafNum]TxConstraints, hFunc mimc.MiMC) error {
	for i, tx := range txs {
		api.AssertIsEqual(TxRootHash, proofs[i].RootHash)

		// index of the leaf
		proofs[i].VerifyProof(api, &hFunc, i)
		api.AssertIsEqual(proofs[i].Path[0], tx.TxHash)

		hFunc.Reset()
		hFunc.Write(tx.TxHash, tx.Timestamp, tx.Nonce)
		hash := hFunc.Sum()
		curve, err := twistededwards.NewEdCurve(api, tedwards.BN254)
		if err != nil {
			return err
		}

		err = sigeddsa.Verify(curve, tx.Sgin, hash, tx.PubKey, &hFunc)
		if err != nil {
			return err
		}
	}

	return nil
}
