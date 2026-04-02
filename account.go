package main

type Account struct {
	Address string
	Balance float64
}

type AccountManager struct {
	accounts map[string]*Account
}

func NewAccountManager() *AccountManager {
	return &AccountManager{accounts: make(map[string]*Account)}
}

func (am *AccountManager) CreateAccount() *Account {
	addr := GenerateAddress()
	acc := &Account{Address: addr, Balance: 0}
	am.accounts[addr] = acc
	return acc
}
