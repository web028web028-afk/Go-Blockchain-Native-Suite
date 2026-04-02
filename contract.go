package main

type Contract struct {
	Address string
	Storage map[string]string
}

type ContractManager struct {
	contracts map[string]*Contract
}

func NewContractManager() *ContractManager {
	return &ContractManager{contracts: make(map[string]*Contract)}
}

func (cm *ContractManager) Deploy() *Contract {
	addr := GenerateAddress()
	c := &Contract{Address: addr, Storage: make(map[string]string)}
	cm.contracts[addr] = c
	return c
}
