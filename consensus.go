package main

func ResolveConflicts(chains []Blockchain) Blockchain {
	longest := chains[0]
	for _, c := range chains {
		if len(c.Blocks) > len(longest.Blocks) && IsChainValid(c) {
			longest = c
		}
	}
	return longest
}
