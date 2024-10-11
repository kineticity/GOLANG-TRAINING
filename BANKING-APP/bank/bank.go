package bank

import (
	"errors"
	"fmt"
)

var banks []Bank
var bankid int = 1

type Bank struct {
	BankID      int
	FullName    string
	Abbreviation string
	IsActive    bool
}

func NewBank(fullName, abbreviation string) (*Bank, error) {  //CREATE

	if fullName == "" {
		return nil, errors.New("full name cannot be empty")
	}
	if abbreviation == "" {
		return nil, errors.New("abbreviation cannot be empty")
	}
	
	b:=&Bank{
		BankID:      bankid,
		FullName:    fullName,
		Abbreviation: abbreviation,
		IsActive:    true, // Active by default
	}

	bankid++
	banks=append(banks, *b)

	return b,nil
}


func GetAllBanks() []*Bank { //GETALL
	var activeBanks []*Bank
	for _, b := range banks {
		if b.IsActive {
			activeBanks = append(activeBanks, &b)
		}
	}
	return activeBanks
}


func GetBankByID(bankID int) (*Bank, error) { //GETBYID
	for i := range banks {
		if banks[i].BankID == bankID && banks[i].IsActive {
			return &banks[i], nil
		}
	}
	return nil, errors.New("bank not found")
}

func DeleteBankByID(bankID int) error { //DELETE
	for i := range banks {
		if banks[i].BankID == bankID {
			banks[i].IsActive = false // softdelete
			fmt.Printf("Bank with ID %d deleted (isActive=false)\n", bankID)
			return nil
		}
	}
	return errors.New("bank not found")
}

func (b *Bank) UpdateBankField(param, newValue string) error { //UPDATE
	if !b.IsActive {
		return errors.New("bank is not active, cannot update")
	}
	if param==""{
		return errors.New("parameter name can't be empty")
	}

	switch param {
	case "bankName":
		if newValue == "" {
			return errors.New("bank name cannot be empty")
		}
		b.FullName = newValue
		fmt.Printf("Bank name updated to: %s\n", b.FullName)
	case "abbreviation":
		if newValue == "" {
			return errors.New("abbreviation cannot be empty")
		}
		b.Abbreviation = newValue
		fmt.Printf("Bank abbreviation updated to: %s\n", b.Abbreviation)
	default:
		return errors.New("invalid parameter name")
	}

	return nil
}
