package account

import (
	"errors"
	"fmt"
	"time"
)

type Transaction struct {
	category  string // credit or debit
	amount    float64
	balance   float64
	accountid int
	bankid    int
	timestamp     time.Time
}

func NewTransaction(category string, amount float64, balance float64, accountid int, bankid int, transactionTime time.Time) (*Transaction, error) {
	if category != "credit" && category != "debit" {
		return nil, errors.New("invalid category: must be 'credit' or 'debit'")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if accountid <= 0 || bankid <= 0 {
		return nil, errors.New("invalid account ID or bank ID")
	}

	transaction := &Transaction{
		category:  category,
		amount:    amount,
		balance:   balance,
		accountid: accountid,
		bankid:    bankid,
		timestamp:      transactionTime,
	}

	return transaction, nil
}


//Getter Setter fns
func (t *Transaction) GetCategory() string {
	return t.category
}

func (t *Transaction) SetCategory(category string) error {
	if category != "credit" && category != "debit" {
		return errors.New("invalid category: must be 'credit' or 'debit'")
	}
	t.category = category
	return nil
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}

func (t *Transaction) SetAmount(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	t.amount = amount
	return nil
}

func (t *Transaction) GetBalance() float64 {
	return t.balance
}

func (t *Transaction) SetBalance(balance float64) {
	t.balance = balance
}

func (t *Transaction) GetAccountID() int {
	return t.accountid
}

func (t *Transaction) SetAccountID(accountid int) error {
	if accountid <= 0 {
		return errors.New("invalid account ID")
	}
	t.accountid = accountid
	return nil
}

func (t *Transaction) GetBankID() int {
	return t.bankid
}

func (t *Transaction) SetBankID(bankid int) error {
	if bankid <= 0 {
		return errors.New("invalid bank ID")
	}
	t.bankid = bankid
	return nil
}

func (t *Transaction) GetTime() time.Time {
	return t.timestamp
}

func (t *Transaction) SetTime(transactionTime time.Time) {
	t.timestamp = transactionTime
}

func (t *Transaction) PrintTransaction() {
	fmt.Println("Transaction Details:")
	fmt.Printf("Category: %s\n", t.category)
	fmt.Printf("Amount: %.2f\n", t.amount)
	fmt.Printf("Balance: %.2f\n", t.balance)
	fmt.Printf("Account ID: %d\n", t.accountid)
	fmt.Printf("Bank ID: %d\n", t.bankid)
	fmt.Printf("Time: %s\n", t.timestamp.Format(time.RFC1123))
	fmt.Println("------------------------------")
}
