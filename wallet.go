package main

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAddress() string {
	b := make([]byte, 20)
	rand.Read(b)
	return "0x" + hex.EncodeToString(b)
}
