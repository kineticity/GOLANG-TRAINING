package services

import (
	"bankingApp/models"
	"errors"
)

var allBanks = []*models.Bank{{BankID: 1, FullName: "SBI", Abbreviation: "SBI", IsActive: true}}
var bankIDCounter = 2

func CreateBank(fullname, abbreviation string) (*models.Bank, error) {
	for _, bank := range allBanks {
		if bank.FullName == fullname {
			return nil, errors.New("bank with same name already exists")
		}
	}

	bank := &models.Bank{
		BankID:      bankIDCounter,
		FullName:    fullname,
		Abbreviation: abbreviation,
		IsActive:    true,
		

	}

	bankIDCounter++
	allBanks = append(allBanks, bank)
	return bank, nil
}

func GetBankByID(bankID int) (*models.Bank, error) {
	for _, bank := range allBanks {
		if bank.BankID == bankID {
			return bank, nil
		}
	}
	return nil, errors.New("bank not found")
}

func GetAllBanks() []*models.Bank {
	return allBanks
}

func UpdateBank(bankID int, fullname, abbreviation string) (*models.Bank, error) {
	for i, bank := range allBanks {
		if bank.BankID == bankID {
			allBanks[i].FullName = fullname
			allBanks[i].Abbreviation = abbreviation
			return allBanks[i], nil
		}
	}
	return nil, errors.New("bank not found")
}

func DeleteBank(bankID int) error {
	for i, bank := range allBanks {
		if bank.BankID == bankID {
			allBanks = append(allBanks[:i], allBanks[i+1:]...)
			return nil
		}
	}
	return errors.New("bank not found")
}



