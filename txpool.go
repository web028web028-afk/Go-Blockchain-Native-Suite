package main

type TxPool struct {
	pending []Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{}
}

func (tp *TxPool) AddTx(tx Transaction) {
	tp.pending = append(tp.pending, tx)
}

func (tp *TxPool) GetPending() []Transaction {
	return tp.pending
}

func (tp *TxPool) Clear() {
	tp.pending = nil
}
