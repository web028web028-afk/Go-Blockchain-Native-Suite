package main

type Token struct {
	Name     string
	Symbol   string
	Total    float64
	Balances map[string]float64
}

func NewToken() *Token {
	return &Token{
		Name:     "GoChainToken",
		Symbol:   "GCT",
		Total:    21000000,
		Balances: make(map[string]float64),
	}
}

func (t *Token) Transfer(from, to string, amount float64) bool {
	if t.Balances[from] < amount {
		return false
	}
	t.Balances[from] -= amount
	t.Balances[to] += amount
	return true
}
