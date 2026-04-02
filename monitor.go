package main

import "fmt"

func ShowChainInfo(chain Blockchain) {
	fmt.Println("Chain Height:", len(chain.Blocks))
	fmt.Println("Latest Hash:", chain.Blocks[len(chain.Blocks)-1].Hash)
}

func IsFork(a, b Blockchain) bool {
	min := len(a.Blocks)
	if len(b.Blocks) < min {
		min = len(b.Blocks)
	}
	for i := 0; i < min; i++ {
		if a.Blocks[i].Hash != b.Blocks[i].Hash {
			return true
		}
	}
	return false
}
