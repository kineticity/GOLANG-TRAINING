package account

import (
	"fmt"
	"time"
)

var allTransactions []*Transaction
type Passbook struct {
	transactions []*Transaction //iska bhi interface?
}

func NewPassbook(initialBalance float64,accountid int,bankid int) (*Passbook,error) {
	transaction, err := NewTransaction("credit", initialBalance, initialBalance, accountid,bankid,time.Now())
	if err!=nil{
		return nil,err
	}

	allTransactions=append(allTransactions, transaction)
	return &Passbook{
		transactions: allTransactions, 
	},nil
}

// Getter Setter fns
func (p *Passbook) GetTransactions() []*Transaction {
	return p.transactions
}

func (p *Passbook) SetTransactions(transactions []*Transaction) {
	p.transactions = transactions
}

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
