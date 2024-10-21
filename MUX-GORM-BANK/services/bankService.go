package services

import (
	database "bankingApp/databases"
	"bankingApp/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db *gorm.DB

// func SetDB(database *gorm.DB) {
// 	db = database
// }

func CreateBank(fullName, abbreviation string) (*models.Bank, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	bank := &models.Bank{
		FullName:     fullName,
		Abbreviation: abbreviation,
		Accounts: []*models.Account{},
		Ledger: []*models.LedgerEntry{},
	}

	err := bank.Create(tx)
	if err != nil {
		return nil, err
	}
	tx.Commit()
	return bank, nil
}

func DeleteBankByID(id int) error {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	bank := &models.Bank{Model: gorm.Model{ID: uint(id)}}
	if err := bank.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateBankByID(id int, fullName, abbreviation string) (*models.Bank, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() 
		}
	}()

	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	bank, err := models.GetBankByID(tx,id)
	if err != nil {
		tx.Rollback() 
		return nil, err
	}

	bank.FullName = fullName
	bank.Abbreviation = abbreviation

	if err := bank.Update(tx); err != nil {
		tx.Rollback() 
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return bank, nil
}

func GetAllBanks() ([]models.Bank, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	banks, err := models.GetAllBanks(tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return banks, nil
}

func GetBankByID(bankID uint) (*models.Bank, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var bank *models.Bank
	bank, err := models.GetBankByID(tx, int(bankID))
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return bank, nil
}
