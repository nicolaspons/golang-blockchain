package blockchain

// BlackChain object that contains all the Block objects.
type BlockChain struct {
	Blocks []*Block
}

// Initialises the BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// Adds a Block object to the a BlockChain object.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}
