package main

import (
	"fmt"
	"math/big"
)

func main() {
	/*chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}*/
	target := big.NewInt(1)
	fmt.Printf("%d\n", target)
	target.Lsh(target, uint(256-14)) // shifts the target
	fmt.Printf("%d\n", target)

}
