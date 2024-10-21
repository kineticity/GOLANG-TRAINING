package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Bank struct {
	gorm.Model
	FullName     string         `json:"full_name"`
	Abbreviation string         `json:"abbreviation"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	Accounts     []*Account     `gorm:"foreignKey:BankID;references:ID"`
	Ledger       []*LedgerEntry `gorm:"foreignKey:BankID;references:ID;foreignKey:CorrespondingBankID;references:ID" json:"ledger"`
}

func (b *Bank) Create(db *gorm.DB) error {
	fmt.Println(b)
	return db.Create(b).Error
}

func (b *Bank) Delete(db *gorm.DB) error {
	return db.Delete(b).Error
}

func GetBankByID(db *gorm.DB, id int) (*Bank, error) {
	var bank Bank
	if err := db.First(&bank, id).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}

func GetAllBanks(db *gorm.DB) ([]Bank, error) {
	var banks []Bank
	if err := db.Find(&banks).Error; err != nil {
		return nil, err
	}
	return banks, nil
}

func (b *Bank) Update(db *gorm.DB) error {
	return db.Save(b).Error
}

func (b *Bank) AddAccount(db *gorm.DB, account *Account) error {
    account.BankID = b.ID



    if err := db.Model(b).Association("Accounts").Append(account).Error; err != nil {
        return err
    }

    return nil
}
