package main

import (
	tedwards "github.com/consensys/gnark-crypto/ecc/twistededwards"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/twistededwards"
	"github.com/consensys/gnark/std/hash/mimc"
	sigeddsa "github.com/consensys/gnark/std/signature/eddsa"
)

type Circuit struct {
	Block              BlockConstraints
	ValidatorSignature sigeddsa.Signature
	ValidatorPubKey    sigeddsa.PublicKey
	BlockHash          frontend.Variable `gnark:",public"`
}

func (c *Circuit) Define(api frontend.API) error {
	hFunc, err := mimc.NewMiMC(api)
	if err != nil {
		return err
	}

	blockHash, err := verify(api, c.Block, hFunc)
	api.AssertIsEqual(c.BlockHash, blockHash)

	curve, err := twistededwards.NewEdCurve(api, tedwards.BN254)
	if err != nil {
		return err
	}

	err = sigeddsa.Verify(curve, c.ValidatorSignature, blockHash, c.ValidatorPubKey, &hFunc)
	return err
}

type BlockConstraints struct {
	BlockNumber   frontend.Variable
	BlockPreImage frontend.Variable
	PreBlockHash  frontend.Variable
	Timestamp     frontend.Variable
}

func verify(api frontend.API, t BlockConstraints, hFunc mimc.MiMC) (frontend.Variable, error) {
	hFunc.Reset()
	hFunc.Write(t.BlockNumber, t.BlockPreImage, t.PreBlockHash, t.Timestamp)
	hash := hFunc.Sum()
	return hash, nil
}
