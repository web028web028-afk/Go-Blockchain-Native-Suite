package main

func IsChainValid(chain Blockchain) bool {
	for i := 1; i < len(chain.Blocks); i++ {
		curr := chain.Blocks[i]
		prev := chain.Blocks[i-1]
		if curr.Hash != CalculateHash(curr) {
			return false
		}
		if curr.PreviousHash != prev.Hash {
			return false
		}
	}
	return true
}
