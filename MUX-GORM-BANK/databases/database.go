package database

import (
	"bankingApp/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	DB, err = gorm.Open("mysql", "root:Forcepointpassword@1@/bankingappnewdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Could not connect to mysql db: %v", err)
	}
	fmt.Println("Connected to the database successfully")

	DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Bank{}, &models.LedgerEntry{}, &models.Transaction{})

	//Manually adding foreign keys
	DB.Model(&models.Account{}).AddForeignKey("customer_id", "users(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Account{}).AddForeignKey("bank_id", "banks(id)", "CASCADE", "CASCADE")
	DB.Model(&models.LedgerEntry{}).AddForeignKey("bank_id", "banks(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Transaction{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "CASCADE")
	DB.Model(&models.LedgerEntry{}).AddForeignKey("corresponding_bank_id", "banks(id)", "CASCADE", "CASCADE")
	DB.Model(&models.Transaction{}).AddForeignKey("corresponding_account_id", "accounts(id)", "CASCADE", "SET NULL")

	return DB

}

func GetDB() *gorm.DB {
	return DB
}
