package main

func AvgBlockTime(chain Blockchain) int64 {
	if len(chain.Blocks) <= 1 {
		return 0
	}
	var total int64
	for i := 1; i < len(chain.Blocks); i++ {
		total += chain.Blocks[i].Timestamp - chain.Blocks[i-1].Timestamp
	}
	return total / int64(len(chain.Blocks)-1)
}
