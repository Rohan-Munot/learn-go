// Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customer's balance. If the customer does not have enough money, it should return the error insufficient funds and leave the balance unchanged. If the transactionType isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil

package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(c *customer, t transaction) error {
	switch t.transactionType {
	case transactionDeposit:
		c.balance += t.amount
	case transactionWithdrawal:
		if c.balance < t.amount {
			return errors.New("insufficient funds")
		}
		c.balance -= t.amount
	default:
		return errors.New("unknown transaction type")
	}
	return nil
}
