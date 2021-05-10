package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// For now, the difficulty is static but it will be computed
// through an algorithm soon.
const Difficulty = 18

// ProofOfWork object used to write a block in the blockchain.
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// Produces a pointer to a ProofOfWork object from
// a pointer to a BLock object.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty)) // shifts the target
	pow := &ProofOfWork{b, target}
	return pow
}

// Takes the data from the block.
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// Creates a hash of the data plus the counter.
// Checks the hash to see if it meets a set of requirements.
// Creates a counter (nonce) which starts at 0.
// Requirements:
// The first few bytes must contain 0s.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

// Tool function that convert an int64 to an array of bytes.
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
