package blockchain

import (
	"bytes"
	"os"
	"testing"
)

func TestValidChain(t *testing.T) {
	blkchain := New()
	blkchain.Add([]byte("NewSTring"))
	blkchain.Add([]byte("NewSTring"))
	err := blkchain.Validate()
	if err != nil {
		t.Error("Could not validate blockchain")
	}
}

//TestGetValidChain creates two different blockchains of different sizes and attempts to validate the chain.
func TestGetValidChain(t *testing.T) {

	blkchain := New()
	blkchain.Add([]byte("NewSTring"))
	blkchain.Add([]byte("NewSTring"))
	err := blkchain.Validate()
	if err != nil {
		t.Error("Failed to Valiate Blockchain")
	}
	blkchain2 := New()
	blkchain2.Add([]byte("chain2str"))
	blkchain2.Add([]byte("chain2str"))
	blkchain2.Add([]byte("data"))
	err = blkchain2.Validate()
	if err != nil {
		t.Error("Failed to Valiate Blockchain")
	}

	//validChain := GetValidChain(blkchain, blkchain2)

	if !bytes.Equal(blkchain[0].Data, blkchain2[0].Data) {
		t.Error("Valid block was not returned")
	}

	if Equal(blkchain, blkchain2) {
		t.Error("Validity check for blockchain binaries failed")
	}
}

//TestBinaryConverter checks for proper encoding and decoding of blockchain gobs to buffer
func TestBinaryConverter(t *testing.T) {
	blkchain := New()
	blkchain.Add([]byte("NewSTring"))
	blkchain.Add([]byte("NewSTring"))
	_ = blkchain.Validate()
	bin := blkchain.encodeChain()

	var out Blockchain
	err := out.decodeChain(bin)
	if err != nil {
		t.Error("Invalid decode Blockchain", err)
	}
}

func TestEqual(t *testing.T) {
	//construct two chains with genesis blocks
	chainA := New()
	var chainB Blockchain

	//Test with equal blocks
	chainA.Add([]byte("NewSTring"))
	chainA.Add([]byte("NewSTring"))

	chainB = chainA

	//Test equality with two equal blockchains
	if !Equal(chainA, chainB) {
		t.Error("equivilent chains are not equal")
	}

	//Test equality with additional block
	chainB.Add([]byte("data"))
	if Equal(chainA, chainB) {
		t.Error("unequivilent chains are testing as equivilent")
	}

}

func TestInputOutput(t *testing.T) {
	//Test saving to file
	blkchain := New()
	blkchain.Add([]byte("NewSTring"))
	blkchain.Add([]byte("NewSTring"))
	err := blkchain.Save("testfile")
	if err != nil {
		t.Error("failed to save blockchain ", err)
	}

	//Test loading from file
	var blkchainB Blockchain
	err = blkchainB.Load("testfile")
	if err != nil {
		t.Error("failed to load blockchain", err)
	}

	//Test validity of read chain
	if !Equal(blkchain, blkchainB) {
		t.Error("read blockchain does not match saved chain")
	}

	//Cleanup test file
	err = os.Remove("testfile.dat")
	if err != nil {
		t.Error("failed to cleanup IO test file")
	}
}
