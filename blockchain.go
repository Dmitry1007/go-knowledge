package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	Blocks []Block
}

// Block represents a block in a Blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	prevHash []byte
}

func (b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.Data, b.prevHash}, []byte{})
	hashed := sha256.Sum256(info)
	b.Hash = hashed[:]
}

func createBlock(data string) Block {
	block := Block{
		Hash:     []byte{},
		Data:     []byte(data),
		prevHash: []byte{},
	}
	block.deriveHash()

	return block
}

func initBlockchain() Blockchain {
	genesis := createBlock("Genesis")
	return Blockchain{Blocks: []Block{genesis}}

}

func (chain *Blockchain) addBlock(data string) Block {
	newBlock := createBlock("data")
	newBlock.prevHash = chain.Blocks[len(chain.Blocks)-1].Hash
	chain.Blocks = append(chain.Blocks, newBlock)
	return newBlock
}

func main() {
	chain := initBlockchain()
	chain.addBlock("block1")
	chain.addBlock("block2")

	for _, block := range chain.Blocks {
		fmt.Println("----------------------------------")
		fmt.Printf("Previous Hash: %x\n", block.prevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
