package bank

import (
	"bankingApp/account"
	"bankingApp/validations"
	"errors"
	"fmt"
	
)

var banks []Bank
var bankid int = 1
type BankOperations interface{
	UpdateBankField(param, newValue string) error
	AddAccount(acc *account.Account)
	GetLedger() Ledger //abstract ledger interface not concrete bankledger struct
	AddTransactionToLedger(otherBank string, amount float64)
	PrintBankLedger()
	LendTo(otherBankname string, amount float64,otherBankid int) error
	ReceiveFrom(otherBank string, amount float64,otherBankid int) error
	Read() string
	Delete()
	GetFullName() string
	GetBankID() int
	

}
type Bank struct {
	bankID       int
	fullName     string
	abbreviation string
	isActive     bool
	accounts     []*account.Account
	ledger Ledger
}

func NewBank(fullName, abbreviation string) (*Bank, error) {

	if err := validation.ValidateNonEmptyString("Full name", fullName); err != nil {
		return nil, err
	}

	if err := validation.ValidateNonEmptyString("Abbreviation", abbreviation); err != nil {
		return nil, err
	}
	tempBankLedger:=NewBankLedger()

	b := &Bank{
		bankID:       bankid,
		fullName:     fullName,
		abbreviation: abbreviation,
		isActive:     true,
		accounts:     []*account.Account{},
		ledger:tempBankLedger,
	}

	bankid++
	banks = append(banks, *b)

	return b, nil
}


// Getter Setter fns
func (b *Bank) GetBankID() int {
	return b.bankID
}

func (b *Bank) SetBankID(id int) {
	b.bankID = id
}

func (b *Bank) GetFullName() string {
	return b.fullName
}

func (b *Bank) SetFullName(name string) {
	b.fullName = name
}

func (b *Bank) GetAbbreviation() string {
	return b.abbreviation
}

func (b *Bank) SetAbbreviation(abbreviation string) {
	b.abbreviation = abbreviation
}

func (b *Bank) GetIsActive() bool {
	return b.isActive
}

func (b *Bank) SetIsActive(active bool) {
	b.isActive = active
}

func GetAllBanks() []BankOperations { //GETALL
	var activeBanks []BankOperations
	for _, b := range banks {
		if b.GetIsActive() {
			activeBanks = append(activeBanks, &b)
		}
	}
	return activeBanks
}

func GetBankByID(bankID int) (BankOperations, error) { //GETBYID
	for i := range banks {
		if banks[i].GetBankID() == bankID && banks[i].GetIsActive() {
			return &banks[i], nil
		}
	}
	return nil, errors.New("bank not found")
}


func DeleteBankByID(bankID int) error { 
	for i := range banks {
		if banks[i].GetBankID() == bankID {
			banks[i].SetIsActive(false)
			fmt.Printf("Bank with ID %d deleted (isActive=false)\n", bankID)
			return nil
		}
	}
	return errors.New("bank not found")
}
func (b *Bank) Read() string {
	return fmt.Sprintf("Bank ID: %d\nFull Name: %s\nAbbreviation: %s\nActive: %t\n", 
		b.GetBankID(), b.GetFullName(), b.GetAbbreviation(), b.GetIsActive())
}
func (b *Bank) UpdateBankField(param, newValue string) error { //UPDATE
	if !b.GetIsActive() {
		return errors.New("bank is not active, cannot update")
	}

	if err := validation.ValidateNonEmptyString("Parameter name", param); err != nil {
		return err
	}
	switch param {
	case "bankName":

		if err := validation.ValidateNonEmptyString("Bank name", newValue); err != nil {
			return err
		}
		b.SetFullName(newValue)
		fmt.Printf("Bank name updated to: %s\n", b.GetFullName())
	case "abbreviation":

		if err := validation.ValidateNonEmptyString("Abbreviation", newValue); err != nil {
			return err
		}
		b.SetAbbreviation(newValue)
		fmt.Printf("Bank abbreviation updated to: %s\n", b.GetAbbreviation())
	default:
		return errors.New("invalid parameter name")
	}

	return nil
}

func (b *Bank) Delete() { //DELETE
		b.SetIsActive(false)	
	}

func (b *Bank) AddAccount(acc *account.Account) {
	b.accounts = append(b.accounts, acc)
	fmt.Printf("Account with ID %d added to Bank ID %d\n", acc.GetAccountID(), b.GetBankID())
}


func (b *Bank) GetLedger() Ledger { //abstract ledger

	return b.ledger
}

func (b *Bank) AddTransactionToLedger(otherBank string, amount float64) {
	b.ledger.AddEntry(otherBank, amount)
}

func (b *Bank) PrintBankLedger() {
	b.ledger.PrintLedger()
}

func (b *Bank) LendTo(otherBankname string, amount float64,otherBankid int) error {
	b.AddTransactionToLedger(otherBankname, -amount) // Amount owed to the other bank
	otherbank,err:=GetBankByID(otherBankid)
	if err!=nil{
		return err
	}
	otherbank.AddTransactionToLedger(b.GetFullName(),+amount)
	return nil
}

func (b *Bank) ReceiveFrom(otherBank string, amount float64,otherBankid int) error {
	b.AddTransactionToLedger(otherBank, +amount) // Amount owed from the other bank
	otherbank,err:=GetBankByID(otherBankid)
	if err!=nil{
		return err
	}
	otherbank.AddTransactionToLedger(b.GetFullName(),-amount)
	return nil
}

