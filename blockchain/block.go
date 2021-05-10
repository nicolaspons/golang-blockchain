package blockchain

// BlackChain object that contains all the Block objects.
type BlockChain struct {
	Blocks []*Block
}

// Block object that contains byte represntation of the data,
// the hash of the previous Block object and its current hash.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Creates a new Block object and fill it with a data string and the
// previous hash.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// Adds a Block object to the a BlockChain object.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// Creates the first BlockChain's Block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Initialises the BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
