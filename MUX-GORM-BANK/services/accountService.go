package services

import (
	database "bankingApp/databases"
	"bankingApp/models"
	"errors"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func CreateAccount(customerID, bankID uint, initialBalance float64) (*models.Account, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	_, err := models.GetBankByID(tx, int(bankID)) 
	if err != nil {
		return nil, err
	}

	err=models.CheckIfAccountExists(tx,customerID,bankID);if err==nil{
		return nil, errors.New("account already exists in this bank")
	}

	account := &models.Account{
		CustomerID: customerID,
		BankID:     bankID,
		Balance:    initialBalance,
	}

	if err := account.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return account, nil
}

func DeleteAccountByID(id int) error {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	account := &models.Account{Model: gorm.Model{ID: uint(id)}}
	if err := account.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetAccountByID(customerID, accountID uint) (*models.Account, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()


	var account *models.Account

	account,err:=models.GetAccountByIDAndCustomerID(tx,accountID,customerID)
	if err!=nil{
		tx.Rollback()
		return nil,err
	}

	return account, tx.Commit().Error
}

func GetAllAccountsByUserID(customerID uint) ([]models.Account, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var accounts []models.Account
	accounts,err:=models.GetAccountsByCustomerID(tx,customerID)
	if err!=nil{
		tx.Rollback()
		return nil,err
	}

	return accounts, tx.Commit().Error
}

func UpdateAccountByID(customerID, accountID uint, updatedAccount models.Account) error {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	_, err := models.GetBankByID(tx, int(updatedAccount.BankID)) //validate if bank exists
	if err != nil {
		return err
	}

	accountexists:=models.CheckIfAccountExists(tx,customerID,updatedAccount.BankID);if accountexists!=nil{
		tx.Rollback()
		return errors.New("account already exists in thsi bank ")
	}


	var account *models.Account

	account,err=models.CheckIfAccountBelongsToCustomer(tx,accountID,customerID)
	if err!=nil{
		tx.Rollback()
		return err
	}

	account.Balance = updatedAccount.Balance
	account.BankID = updatedAccount.BankID

	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func Withdraw(id int, amount float64) (*models.Account, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	account, err := models.GetAccountByID(tx, id)
	if err != nil {
		return nil, err
	}

	if err := account.Withdraw(tx, amount); err != nil {
		tx.Rollback()
		return nil, err
	}

	transaction := &models.Transaction{
		TransactionType:        "withdraw",
		Amount:                 amount,
		Time:                   time.Now(),
		NewBalance:             account.Balance,
		CorrespondingAccountID: account.ID,
		AccountID:              account.ID,
	}

	if err := transaction.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := account.AddToPassbook(tx, transaction); err != nil {
		tx.Rollback()
		return nil, err
	}

	return account, tx.Commit().Error
}

func Deposit(id int, amount float64) (*models.Account, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	account, err := models.GetAccountByID(tx, id)
	if err != nil {
		return nil, err
	}

	if err := account.Deposit(tx, amount); err != nil {
		tx.Rollback()
		return nil, err
	}

	transaction := &models.Transaction{
		TransactionType:        "deposit",
		Amount:                 amount,
		Time:                   time.Now(),
		NewBalance:             account.Balance,
		CorrespondingAccountID: account.ID,
		AccountID:              account.ID,
	}

	if err := transaction.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := account.AddToPassbook(tx, transaction); err != nil {
		tx.Rollback()
		return nil, err
	}

	return account, tx.Commit().Error
}

func Transfer(fromAccountID, toAccountID int, amount float64) error {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	fromAccount, err := models.GetAccountByID(tx, fromAccountID)
	if err != nil {
		return err
	}

	toAccount, err := models.GetAccountByID(tx, toAccountID)
	if err != nil {
		return err
	}

	if err := fromAccount.Transfer(tx, toAccount, amount); err != nil {
		tx.Rollback()
		return err
	}

	transactionFrom := &models.Transaction{
		TransactionType:        "transfer",
		Amount:                 -amount,
		Time:                   time.Now(),
		NewBalance:             fromAccount.Balance,
		CorrespondingAccountID: toAccount.ID,
		AccountID:              fromAccount.ID,
	}

	if err := transactionFrom.Create(tx); err != nil {
		tx.Rollback()
		return err
	}

	transactionTo := &models.Transaction{
		TransactionType:        "transfer",
		Amount:                 amount,
		Time:                   time.Now(),
		NewBalance:             toAccount.Balance,
		CorrespondingAccountID: fromAccount.ID,
		AccountID:              toAccount.ID,
	}

	if err := transactionTo.Create(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := toAccount.AddToPassbook(tx, transactionTo); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func PrintAccountPassbook(accountID int) error {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	account, err := models.GetAccountByID(tx, accountID) // Fetch account by ID
	if err != nil {
		return err
	}

	// Print the account's passbook
	err = account.PrintPassbook(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

