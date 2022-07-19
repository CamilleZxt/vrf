package vrf

import (
	"fmt"
	"math/big"
	"vrf/common/signatures/secp256k1"
	"vrf/proof"
	"vrf/vrfkey"
)

var (
	rawSecretKey = big.NewInt(1) // never do this in production!
	secretKey    = vrfkey.MustNewV2XXXTestingOnly(rawSecretKey)
	publicKey    = (&secp256k1.Secp256k1{}).Point().Mul(secp256k1.IntToScalar(
		rawSecretKey), nil)
	hardcodedSeed = big.NewInt(0)
	vrfFee        = big.NewInt(7)
)

func GenerateProof() vrfkey.Proof {
	preSeed, err := proof.BigToSeed(big.NewInt(1))
	if err != nil {
		fmt.Println("err:", err)
	}
	s := proof.PreSeedData{
		PreSeed: preSeed,
		//BlockHash: log.Raw.Raw.BlockHash,
		//BlockNum:  log.Raw.Raw.BlockNumber,
	}
	seed := proof.FinalSeed(s)
	proof, err := secretKey.GenerateProofWithNonce(seed, big.NewInt(1) /* nonce */)
	if err != nil {
		fmt.Println("err:", err)
	}
	//proofBlob, err := GenerateProofResponseFromProof(proof, s)
	return proof
}

func GenerateProofResponseFromProof(p vrfkey.Proof, s proof.PreSeedData) (
	proof.MarshaledOnChainResponse, error) {
	return proof.GenerateProofResponseFromProof(p, s)
}
