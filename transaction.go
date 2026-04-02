package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Transaction struct {
	TxID      string
	Sender    string
	Recipient string
	Amount    float64
	Timestamp int64
}

func NewTransaction(sender, recipient string, amount float64) Transaction {
	tx := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
	}
	hash := sha256.Sum256([]byte(sender + recipient + strconv.FormatFloat(amount, 'f', 2, 64) + strconv.FormatInt(tx.Timestamp, 10)))
	tx.TxID = hex.EncodeToString(hash[:])
	return tx
}
