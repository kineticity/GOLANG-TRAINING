package services

import (
	database "bankingApp/databases"
	"bankingApp/models"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateLedgerEntry(bankID, corrBankID uint, amount float64,entrytype string) (*models.LedgerEntry, error) {
	tx := database.GetDB().Begin()
	fmt.Println("transaction starts")
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	fmt.Println("above entry struct")
	var entry1 *models.LedgerEntry
	if entrytype=="lending"{
		entry1 = &models.LedgerEntry{
			BankID:              bankID,
			CorrespondingBankID: corrBankID,
			Amount:              amount,
			EntryType: entrytype,
		}

	}else if entrytype=="receiving"{
		entry1 = &models.LedgerEntry{
			BankID:              bankID,
			CorrespondingBankID: corrBankID,
			Amount:              -amount,
			EntryType: entrytype,
		}

	} else{
		return nil,fmt.Errorf("%v invalid entry type.should be lending or receiving",entrytype)
	}
	
	fmt.Println("below entry struct")


	err := entry1.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating ledger entry: %v", err)
	}
	var entry2 *models.LedgerEntry
	if entrytype=="lending"{
		entry2 = &models.LedgerEntry{
			BankID:              corrBankID,
			CorrespondingBankID: bankID,
			Amount:              amount,
			EntryType: "receiving",
		}

	}else if entrytype=="receiving"{
		entry2 = &models.LedgerEntry{
			BankID:              corrBankID,
			CorrespondingBankID: bankID,
			Amount:              -amount,
			EntryType: "lending",
		}

	} else{
		return nil,fmt.Errorf("invalid entry type.should be lending or receiving")
	}
	
	fmt.Println("below entry struct")


	err = entry2.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating ledger entry: %v", err)
	}

	return nil, tx.Commit().Error

	
}

func GetLedgerEntries(bankID int) ([]models.LedgerEntry, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return models.GetLedgerEntries(tx,bankID)
}
