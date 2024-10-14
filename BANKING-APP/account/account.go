package account

import (
	"errors"
	"time"
	"bankingApp/validations"
)

type Account struct {
	accountID int
	balance   float64
	isActive  bool
	bankid int
	passbook  *Passbook 
}



var allAccounts []*Account
var accountid int = 1

func NewAccount(initialBalance float64,bankid int) (*Account, error) { //bankid is validated in customer's createaccount
	if initialBalance < 1000 {
		return nil, errors.New("initial balance must be at least Rs. 1000")
	}
	passbook,err:=NewPassbook(initialBalance,accountid,bankid)
	if err!=nil{
		return nil,err
	}

	account := &Account{
		accountID: accountid,
		balance:   initialBalance,
		isActive:  true,
		bankid:bankid,
		passbook:  passbook,
	}
	accountid++
	allAccounts = append(allAccounts, account)


	return account, nil
}


// Getter Setter fns
func (a *Account) GetAccountID() int {
	return a.accountID
}

func (a *Account) SetAccountID(id int) {
	a.accountID = id
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) SetBalance(balance float64) {
	a.balance = balance
}

func (a *Account) GetIsActive() bool {
	return a.isActive
}

func (a *Account) SetIsActive(active bool) {
	a.isActive = active
}

func (a *Account) GetBankID() int {
	return a.bankid
}

func (a *Account) SetBankID(bankid int) {
	a.bankid = bankid
}

func (a *Account) GetPassbook() *Passbook {
	return a.passbook
}

func (a *Account) SetPassbook(passbook *Passbook) { //YAGNI
	a.passbook = passbook
}


func GetAccountByID(accountID int,accounts []*Account) (*Account, error) { //GETBYID
	for _, acc := range accounts { //DRY?
		if acc.GetAccountID() == accountID && acc.GetIsActive() {
			return acc, nil
		}
	}
	return nil, errors.New("account not found or inactive")
}

func GetAllAccounts()[]*Account{
	return allAccounts

}


func (a *Account) UpdateAccount(attribute string, newValue interface{}) error { //UPDATE
	switch attribute {
	case "balance":
		if value, ok := newValue.(float64); ok {

			if err := validation.ValidatePositiveNumber("Balance", value); err != nil {
				return err
			}
			a.SetBalance(value)
			return nil
		}
		return errors.New("invalid value for balance")
	case "isActive":
		if value, ok := newValue.(bool); ok {
			if !value && a.GetBalance() > 0 {
				return errors.New("cannot deactivate account with a positive balance")
			}
			a.SetIsActive(value)
			return nil
		}
		return errors.New("invalid value for isActive")
	default:
		return errors.New("unknown attribute")
	}
}

func (a *Account) Deactivate() { //DELETE
	a.SetIsActive(false)
}



func (a *Account) Deposit(amount float64,accountid int,bankid int,transfer bool,fromacc int,frombankid int) error {

	if amount <= 0 { //DRY
		return errors.New("deposit amount must be greater than zero")
	}

	newBalance := a.GetBalance() + amount
	a.SetBalance(newBalance)

	if transfer{ //if the deposit is being called from transfer method of customer
	
		transaction,err:=NewTransaction("credit",amount,newBalance,fromacc,frombankid,time.Now())
		if err!=nil{ 
			return err
		}

		a.passbook.AddTransaction(transaction)

		return nil

	}else{

		transaction,err:=NewTransaction("credit",amount,newBalance,accountid,bankid,time.Now())
		if err!=nil{
			return err
		}
		a.passbook.AddTransaction(transaction)

		return nil
}
}

func (a *Account) Withdraw(amount float64,accountid int,bankid int,transfer bool,toacc int,tobankid int) error {


	if err := validation.ValidatePositiveNumber("Amount", amount); err != nil {
		return err
	}
	if a.GetBalance() < amount {
		return errors.New("insufficient funds")
	}

	newBalance := a.GetBalance() - amount
	a.SetBalance(newBalance)

	if transfer{ //if withdraw was called from transfer method of customer
		transaction,err:=NewTransaction("debit",amount,newBalance,toacc,tobankid,time.Now())
		if err!=nil{
			return err
		}

		a.passbook.AddTransaction(transaction)

		return nil

	}else{
		transaction,err:=NewTransaction("debit",amount,newBalance,accountid,bankid,time.Now())
		if err!=nil{
			return err
		}

		a.passbook.AddTransaction(transaction)

		return nil

	}


}

