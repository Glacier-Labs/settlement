package main

import (
	"crypto/sha256"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
)

var hFunc = mimc.NewMiMC()

// Convert input to the circuit's suitable variable
func sha256Sum(input []byte) []byte {
	s := sha256.New()
	s.Write(input)
	d := s.Sum([]byte{})

	var x fr.Element
	x.SetBigInt(new(big.Int).SetBytes(d))
	d0 := x.Bytes()
	return d0[:]
}
