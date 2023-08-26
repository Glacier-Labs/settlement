package main

import (
	"bytes"
	"fmt"
	"hash"

	"github.com/consensys/gnark-crypto/accumulator/merkletree"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/frontend"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards/eddsa"
)

// A GlacierDB Client for zk-sign test
type BlockCommitment struct {
	BlockNumber   uint64
	BlockPreImage []byte
	PreBlockHash  []byte
	Timestamp     uint64
	PubKey        eddsa.PublicKey
	Signature     eddsa.Signature // signature of the validator
	Txs           []Tx
	Witness       Circuit
}

type Tx struct {
	TxHash    []byte
	Timestamp uint64
	Nonce     uint64
	PubKey    eddsa.PublicKey
	Signature eddsa.Signature // signature of the tx sender
}

func (tx *Tx) Sign(priv eddsa.PrivateKey, h hash.Hash) error {
	tx.PubKey = priv.PublicKey
	h.Reset()

	var frNs fr.Element
	var frTs fr.Element

	// serializing transfer. The signature is on h(nonce, txHash, timestamp)
	// (each pubkey consist of 2 chunks of 256bits)
	frNs.SetUint64(tx.Nonce)
	frTs.SetUint64(tx.Timestamp)

	b := frNs.Bytes()
	_, _ = h.Write(b[:])

	h.Write(tx.TxHash)

	b = frTs.Bytes()
	_, _ = h.Write(b[:])

	hash := h.Sum([]byte{})

	h.Reset()
	sigBin, err := priv.Sign(hash, h)
	if err != nil {
		return err
	}
	var sig eddsa.Signature
	if _, err := sig.SetBytes(sigBin); err != nil {
		return err
	}
	tx.Signature = sig
	return nil
}

// NewTransfer creates a new transfer (to be signed)
func NewBlockCommitment(blockPreImage, preBlockHash []byte, blockNumber, timestamp uint64, pubKey eddsa.PublicKey, txs []Tx) BlockCommitment {
	var res BlockCommitment
	res.BlockNumber = blockNumber
	res.BlockPreImage = blockPreImage
	res.PreBlockHash = preBlockHash
	res.Timestamp = timestamp
	res.PubKey = pubKey
	res.Txs = txs
	return res
}

// Sign signs a transaction
func (t *BlockCommitment) SignWitness(priv eddsa.PrivateKey, h hash.Hash) (eddsa.Signature, error) {
	// txs
	var buf []byte
	for _, tx := range t.Txs {
		buf = append(buf, tx.TxHash[:]...)
	}

	for i, tx := range t.Txs {
		merkleRoot, proofInclusion, numLeaves, err := merkletree.BuildReaderProof(bytes.NewReader(buf), hFunc, hFunc.Size(), uint64(i))
		if err != nil {
			panic(err)
		}

		result := merkletree.VerifyProof(hFunc, merkleRoot, proofInclusion, uint64(i), numLeaves)

		fmt.Println(numLeaves, result, len(proofInclusion))

		t.Witness.TxRootHash = merkleRoot
		t.Witness.TxProofs[i].RootHash = merkleRoot
		t.Witness.TxProofs[i].Path = make([]frontend.Variable, len(proofInclusion))
		for j := 0; j < len(proofInclusion); j++ {
			t.Witness.TxProofs[i].Path[j] = proofInclusion[j]
		}

		t.Witness.Txs[i] = TxConstraints{
			TxHash:    tx.TxHash,
			Timestamp: tx.Timestamp,
			Nonce:     tx.Nonce,
		}
		t.Witness.Txs[i].PubKey.A.X = tx.PubKey.A.X
		t.Witness.Txs[i].PubKey.A.Y = tx.PubKey.A.Y
		t.Witness.Txs[i].Sgin.R.X = tx.Signature.R.X
		t.Witness.Txs[i].Sgin.R.Y = tx.Signature.R.Y
		t.Witness.Txs[i].Sgin.S = tx.Signature.S[:]
	}

	// block
	hash := t.Hash(h)

	sigBin, err := priv.Sign(hash, hFunc)
	if err != nil {
		return eddsa.Signature{}, err
	}
	var sig eddsa.Signature
	if _, err := sig.SetBytes(sigBin); err != nil {
		return eddsa.Signature{}, err
	}
	t.Signature = sig

	ok, err := t.Verify(h)
	if err != nil {
		return eddsa.Signature{}, err
	}
	if !ok {
		return eddsa.Signature{}, fmt.Errorf("verify not match")
	}

	t.Witness.ValidatorPubKey.A.X = t.PubKey.A.X
	t.Witness.ValidatorPubKey.A.Y = t.PubKey.A.Y

	t.Witness.Block.BlockNumber = t.BlockNumber
	t.Witness.Block.BlockPreImage = t.BlockPreImage
	t.Witness.Block.PreBlockHash = t.PreBlockHash
	t.Witness.Block.Timestamp = t.Timestamp

	t.Witness.ValidatorSignature.R.X = t.Signature.R.X
	t.Witness.ValidatorSignature.R.Y = t.Signature.R.Y
	t.Witness.ValidatorSignature.S = t.Signature.S[:]

	t.Witness.BlockHash = hash

	return sig, nil
}

func (t *BlockCommitment) Verify(h hash.Hash) (bool, error) {
	hash := t.Hash(h)

	// verification of the signature
	resSig, err := t.PubKey.Verify(t.Signature.Bytes(), hash, hFunc)
	if err != nil {
		return false, err
	}
	if !resSig {
		return false, fmt.Errorf("sign not match")
	}
	return true, nil
}

func (t *BlockCommitment) Hash(h hash.Hash) []byte {
	h.Reset()
	var frBn fr.Element
	var frTs fr.Element

	// serializing transfer. The signature is on h(blocknumber, blockpreimage, preblockhash, timestamp)
	// (each pubkey consist of 2 chunks of 256bits)
	frBn.SetUint64(t.BlockNumber)
	frTs.SetUint64(t.Timestamp)

	b := frBn.Bytes()
	_, _ = h.Write(b[:])

	h.Write(t.BlockPreImage)

	h.Write(t.PreBlockHash)

	b = frTs.Bytes()
	_, _ = h.Write(b[:])

	return h.Sum([]byte{})
}
