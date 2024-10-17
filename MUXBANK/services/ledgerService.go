package services

import (
	"bankingApp/models"
	"errors"
	"time"
)

func AddToLedger(lendingBankID, receivingBankID int, amount float64) error {
	var lendingBank, receivingBank *models.Bank

	for _, bank := range allBanks {
		if bank.BankID == lendingBankID {
			lendingBank = bank
		}
		if bank.BankID == receivingBankID {
			receivingBank = bank
		}
	}

	if lendingBank == nil || receivingBank == nil {
		return errors.New("one of the banks does not exist")
	}

	lendingBank.Ledger = append(lendingBank.Ledger, models.LedgerEntry{
		BankName: receivingBank.FullName,
		Amount:   -amount,
		Time:     time.Now(),
	})

	receivingBank.Ledger = append(receivingBank.Ledger, models.LedgerEntry{
		BankName: lendingBank.FullName,
		Amount:   amount,
		Time:     time.Now(),
	})

	return nil
}

func GetLedger(bankID int) ([]models.LedgerEntry, error) {
	for _, bank := range allBanks {
		if bank.BankID == bankID {
			return bank.Ledger, nil
		}
	}
	return nil, errors.New("bank not found")
}
