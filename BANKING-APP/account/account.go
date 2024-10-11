package account

import (
	"errors"
)

type Account struct {
	AccountID int
	Balance   float64
	IsActive  bool
}

var allAccounts []*Account
var accountid int=1


func NewAccount(initialBalance float64) (*Account, error) { //accountid is maintained in customer package //CREATE

	if initialBalance < 1000 { 
		return nil, errors.New("initial balance must be at least Rs. 1000")
	}
	account:=&Account{
		AccountID: accountid,
		Balance:   initialBalance,
		IsActive:  true, 
	}
	accountid++
	allAccounts = append(allAccounts, account) 
	return account,nil
}


func GetAccountByID(accountID int,accounts []*Account) (*Account, error) { //GETBYID
	for _, acc := range accounts {
		if acc.AccountID == accountID && acc.IsActive {
			return acc, nil
		}
	}
	return nil, errors.New("account not found or inactive")
}

func GetGlobalAccountByID(accountID int) (*Account, error) { //GETBYID
	for _, acc := range allAccounts {
		if acc.AccountID == accountID && acc.IsActive {
			return acc, nil
		}
	}
	return nil, errors.New("account not found or inactive")
}

func (a *Account) UpdateAccount(attribute string, newValue interface{}) error { //UPDATE
	switch attribute {
	case "balance":
		if value, ok := newValue.(float64); ok {
			if value < 0 {
				return errors.New("balance cannot be negative")
			}
			a.Balance = value
			return nil
		}
		return errors.New("invalid value for balance")
	case "isActive":
		if value, ok := newValue.(bool); ok {
			if !value && a.Balance > 0 {
				return errors.New("cannot deactivate account with a positive balance")
			}
			a.IsActive = value
			return nil
		}
		return errors.New("invalid value for isActive")
	default:
		return errors.New("unknown attribute")
	}
}

// Deactivate sets the account as inactive
func (a *Account) Deactivate() { //DELETE
	a.IsActive = false
}



func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}
	if a.Balance < amount {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}
