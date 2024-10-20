package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LedgerEntry struct {
	gorm.Model
	BankName            string  `json:"bank_name"`
	Amount              float64 `json:"amount"`
	BankID              uint    `json:"bank_id" gorm:"not null"` // Foreign key to Bank
	CorrespondingBankID uint    `json:"corresponding_bank_id"`
	EntryType           string  `json:"entry_type" gorm:"not null"` // Lending or Receiving

}

func (l *LedgerEntry) Create(db *gorm.DB) error {
	return db.Create(l).Error
}

func GetLedgerEntries(db *gorm.DB, bankID int) ([]LedgerEntry, error) {
	var ledgerEntries []LedgerEntry
	if err := db.Where("bank_id = ?", bankID).Find(&ledgerEntries).Error; err != nil {
		return nil, err
	}
	return ledgerEntries, nil
}
