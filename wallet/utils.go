package wallet

import (
	"log"

	"github.com/mr-tron/base58"
)

// Base58 was invented with Bitcoin. It's a derivative of the base 64 algorithm
// except the main difference is that it uses 6 less characters inside its
// alphabet: O 0 1 I + /
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	if err != nil {
		log.Panic(err)
	}

	return decode
}
