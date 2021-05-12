package blockchain

import (
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// BlackChain object that contains all the Block objects.
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// Initialises the BlockChain
func InitBlockChain() *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err := db.Update(func(txn))
}

// Adds a Block object to the a BlockChain object.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
