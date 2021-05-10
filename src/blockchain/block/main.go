package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlackChain object that contains all the Block objects.
type BlockChain struct {
	blocks []*Block
}

// Block object that contains byte represntation of the data,
// the hash of the previous Block object and its current hash.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Computes the sha256 hash of the Block object thanks to its data
// and the previous hash.
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// Creates a new Block object and fill it with a data string and the
// previous hash.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// Adds a Block object to the a BlockChain object.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Creates the first BlockChain's Block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initialises the BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
