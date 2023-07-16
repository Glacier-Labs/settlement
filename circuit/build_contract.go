package main

import (
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// step0: compile circuit
	var circuit Circuit

	r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuit)
	checkError(err)

	pk, vk, err := groth16.Setup(r1cs)
	checkError(err)

	{
		f, err := os.Create("./contract/zkglacier.g16.vk")
		checkError(err)

		_, err = vk.WriteRawTo(f)
		checkError(err)
	}
	{
		f, err := os.Create("./contract/zkglacier.g16.pk")
		checkError(err)

		_, err = pk.WriteRawTo(f)
		checkError(err)
	}

	{
		f, err := os.Create("./contract/zkglacier_g16.sol")
		checkError(err)

		err = vk.ExportSolidity(f)
		checkError(err)
	}
}
