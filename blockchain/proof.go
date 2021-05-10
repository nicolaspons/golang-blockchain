package blockchain

import (
	"bytes"
	"math/big"
)

// For now, the difficulty is static but it will be computed
// through an algorithm soon.
const Difficulty = 12

// Takes the data from the block.
// Creates a counter (nonce) which starts at 0.
// Creates a hash of the data plus the counter.
// Checks the hash to see if it meets a set of requirements.
// Requirements:
// The first few bytes must contain 0s.
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// Produces a pointer to a ProofOfWork object from
// to a pointer to a BLock object.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // shifts the target
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{pow.Block.PrevHash,
			pow.Block.Data,
		},
		[]byte{},
	)
	return data
}
