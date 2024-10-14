package account

import (
	"fmt"
	"time"
)

var allTransactions []*Transaction
type Passbook struct {
	transactions []*Transaction
}

// Factory function to create a new Passbook with a nil transaction
func NewPassbook(initialBalance float64,accountid int,bankid int) (*Passbook,error) {
	transaction, err := NewTransaction("credit", initialBalance, initialBalance, accountid,bankid,time.Now())
	if err!=nil{
		return nil,err
	}

	allTransactions=append(allTransactions, transaction)
	return &Passbook{
		transactions: allTransactions, // Initialize with a nil transaction
	},nil
}

// Getter for transaction
func (p *Passbook) GetTransactions() []*Transaction {
	return p.transactions
}

// Setter for transaction
func (p *Passbook) SetTransactions(transactions []*Transaction) {
	p.transactions = transactions
}

// AddTransaction method to add a new transaction to the passbook
func (p *Passbook) AddTransaction(t *Transaction) {
	p.transactions = append(p.transactions, t)
}

func (p *Passbook) PrintPassbook() {
	if len(p.GetTransactions()) == 0 {
		fmt.Println("No transactions available in the passbook.")
		return
	}

	fmt.Println("Passbook Transactions:")
	for _, transaction := range p.transactions {

		transaction.PrintTransaction()
	}
}
