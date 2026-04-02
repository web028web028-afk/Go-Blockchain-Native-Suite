package main

import (
	"strings"
	"strconv"
)

const difficulty = 4

func ProofOfWork(block Block) int {
	nonce := 0
	target := strings.Repeat("0", difficulty)
	for {
		block.Nonce = nonce
		hash := CalculateHash(block)
		if strings.HasPrefix(hash, target) {
			break
		}
		nonce++
	}
	return nonce
}
