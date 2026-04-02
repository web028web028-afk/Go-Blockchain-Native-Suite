package main

import "crypto/sha256"

func MerkleRoot(hashes [][]byte) []byte {
	if len(hashes) == 0 {
		return nil
	}
	for len(hashes) > 1 {
		if len(hashes)%2 != 0 {
			hashes = append(hashes, hashes[len(hashes)-1])
		}
		var newLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			combined := append(hashes[i], hashes[i+1]...)
			h := sha256.Sum256(combined)
			newLevel = append(newLevel, h[:])
		}
		hashes = newLevel
	}
	return hashes[0]
}
