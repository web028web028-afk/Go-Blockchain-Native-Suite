package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	Timestamp    int64
	Data         string
	Nonce        int
	Hash         string
}

func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.PreviousHash + strconv.FormatInt(block.Timestamp, 10) + block.Data + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}
