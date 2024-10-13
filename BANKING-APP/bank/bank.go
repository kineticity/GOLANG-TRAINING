package bank

import (
	"bankingApp/account"
	"errors"
	"fmt"
	
)

var banks []Bank
var bankid int = 1

type Bank struct {
	bankID       int
	fullName     string
	abbreviation string
	isActive     bool
	accounts     []*account.Account
	ledger *BankLedger
}

func NewBank(fullName, abbreviation string) (*Bank, error) {
	if fullName == "" {
		return nil, errors.New("full name cannot be empty")
	}
	if abbreviation == "" {
		return nil, errors.New("abbreviation cannot be empty")
	}

	b := &Bank{
		bankID:       bankid,
		fullName:     fullName,
		abbreviation: abbreviation,
		isActive:     true,
		accounts:     []*account.Account{},
		ledger:NewBankLedger(),
	}

	bankid++
	banks = append(banks, *b)

	return b, nil
}


// // Bank represents a bank that maintains a ledger with other banks.
// type Bank struct {
// 	bankID   int
// 	fullName string
// 	abbreviation string
// 	accounts map[int]*account.Account
// 	ledger   *ledger.BankLedger // New field for bank ledger
// }

// // NewBank creates a new Bank instance.
// func NewBank(fullName, abbreviation string) (*Bank, error) {
// 	if fullName == "" || abbreviation == "" {
// 		return nil, errors.New("bank name and abbreviation cannot be empty")
// 	}
// 	return &Bank{
// 		fullName:     fullName,
// 		abbreviation: abbreviation,
// 		accounts:     make(map[int]*account.Account),
// 		ledger:       ledger.NewBankLedger(), // Initialize ledger
// 	}, nil
// }

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

func GetAllBanks() []*Bank { //GETALL
	var activeBanks []*Bank
	for _, b := range banks {
		if b.GetIsActive() {
			activeBanks = append(activeBanks, &b)
		}
	}
	return activeBanks
}

func GetBankByID(bankID int) (*Bank, error) { //GETBYID
	for i := range banks {
		if banks[i].GetBankID() == bankID && banks[i].GetIsActive() {
			return &banks[i], nil
		}
	}
	return nil, errors.New("bank not found")
}

func DeleteBankByID(bankID int) error { //DELETE
	for i := range banks {
		if banks[i].GetBankID() == bankID {
			banks[i].SetIsActive(false)
			fmt.Printf("Bank with ID %d deleted (isActive=false)\n", bankID)
			return nil
		}
	}
	return errors.New("bank not found")
}

func (b *Bank) UpdateBankField(param, newValue string) error { //UPDATE
	if !b.GetIsActive() {
		return errors.New("bank is not active, cannot update")
	}
	if param == "" {
		return errors.New("parameter name can't be empty")
	}

	switch param {
	case "bankName":
		if newValue == "" {
			return errors.New("bank name cannot be empty")
		}
		b.SetFullName(newValue)
		fmt.Printf("Bank name updated to: %s\n", b.GetFullName())
	case "abbreviation":
		if newValue == "" {
			return errors.New("abbreviation cannot be empty")
		}
		b.SetAbbreviation(newValue)
		fmt.Printf("Bank abbreviation updated to: %s\n", b.GetAbbreviation())
	default:
		return errors.New("invalid parameter name")
	}

	return nil
}

func (b *Bank) AddAccount(acc *account.Account) {
	b.accounts = append(b.accounts, acc)
	fmt.Printf("Account with ID %d added to Bank ID %d\n", acc.GetAccountID(), b.GetBankID())
}


// GetLedger returns the bank's ledger.
func (b *Bank) GetLedger() *BankLedger {
	return b.ledger
}

// AddTransactionToLedger adds a transaction to the bank's ledger.
func (b *Bank) AddTransactionToLedger(otherBank string, amount float64) {
	b.ledger.AddEntry(otherBank, amount)
}

// PrintBankLedger prints the bank's ledger.
func (b *Bank) PrintBankLedger() {
	b.ledger.PrintLedger()
}

// Example method for a bank to lend money to another bank.
func (b *Bank) LendTo(otherBankname string, amount float64,otherBankid int) error {
	b.AddTransactionToLedger(otherBankname, -amount) // Amount owed to the other bank
	otherbank,err:=GetBankByID(otherBankid)
	if err!=nil{
		return err
	}
	otherbank.AddTransactionToLedger(b.GetFullName(),+amount)
	return nil
}

// Example method for a bank to receive money from another bank.
func (b *Bank) ReceiveFrom(otherBank string, amount float64,otherBankid int) error {
	b.AddTransactionToLedger(otherBank, +amount) // Amount owed from the other bank
	otherbank,err:=GetBankByID(otherBankid)
	if err!=nil{
		return err
	}
	otherbank.AddTransactionToLedger(b.GetFullName(),-amount)
	return nil
}

