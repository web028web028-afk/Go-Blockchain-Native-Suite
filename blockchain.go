package main

import "time"

type Blockchain struct {
	Blocks []Block
}

func InitBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:        0,
		PreviousHash: "0",
		Timestamp:    time.Now().Unix(),
		Data:         "Genesis Block",
		Nonce:        0,
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return &Blockchain{Blocks: []Block{genesisBlock}}
}
