package services

import (
	// "bankingApp/controllers"
	"bankingApp/models"
	"errors"
	"fmt"
	"time"
)

var allAccounts = []*models.Account{}
var accountIDCounter = 1

func CreateAccount(customerID, bankID int, initialBalance float64) (*models.Account, error) {
	// Check if the bank exists
	if !bankExists(bankID) {
		return nil, errors.New("bank does not exist") //bankid validate
	}

	account := &models.Account{
		AccountID:  accountIDCounter,
		CustomerID: customerID,
		BankID:     bankID,
		Balance:    initialBalance,
	}

	accountIDCounter++
	allAccounts = append(allAccounts, account)

	fmt.Println("Is this it????")
	for i, customer := range models.Userslist {
		if customer.UserID == customerID {
			models.Userslist[i].Accounts = append(models.Userslist[i].Accounts, account)
			break
		}
	}

	return account, nil
}

func bankExists(bankID int) bool {
	for _, bank := range allBanks {
		if bank.BankID == bankID {
			return true
		}
	}
	return false
}

func GetAccountByID(accountID int) (*models.Account, error) {
	for _, account := range allAccounts {
		if account.AccountID == accountID {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

func GetAllAccounts(customerID int) []*models.Account {
	for _, customer := range allCustomers {
		if customer.UserID == customerID {
			return customer.Accounts
		}
	}
	return nil

	// return allAccounts
}

func UpdateAccount(accountID int, bankID int, balance float64) (*models.Account, error) {
	for i, account := range allAccounts {
		if account.AccountID == accountID {
			allAccounts[i].BankID = bankID
			allAccounts[i].Balance = balance
			return allAccounts[i], nil
		}
	}
	return nil, errors.New("account not found")
}

func DeleteAccount(accountID int,customerID int) error {
	for i, account := range allAccounts {
		if account.AccountID == accountID {
			allAccounts = append(allAccounts[:i], allAccounts[i+1:]...)
			return nil
		}
	}

	for i, customer := range allCustomers {
		if customer.UserID == customerID {
			customer.Accounts=append(customer.Accounts[:i],customer.Accounts[i+1:]... )
		}
	}
	return errors.New("account not found")
}

func Withdraw(accountID int, amount float64) (*models.Account, error) {
	for _, account := range allAccounts {
		if account.AccountID == accountID {
			if account.Balance < amount {
				return nil, errors.New("insufficient balance")
			}
			account.Balance -= amount
			addTransactionToPassbook(account, "withdraw", amount, account.Balance, 0)
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

func Deposit(accountID int, amount float64) (*models.Account, error) {
	for _, account := range allAccounts {
		if account.AccountID == accountID {
			account.Balance += amount
			addTransactionToPassbook(account, "deposit", amount, account.Balance, 0)
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}

func Transfer(fromAccountID, toAccountID int, amount float64) error {
	var fromAccount, toAccount *models.Account

	if fromAccountID==toAccountID{
		return errors.New("you cannot transfer to the same account")
	}

	for _, account := range allAccounts {
		if account.AccountID == fromAccountID {
			fromAccount = account
		}
		if account.AccountID == toAccountID {
			toAccount = account
		}
	}

	if fromAccount == nil || toAccount == nil {
		return errors.New("one of the accounts does not exist")
	}

	if fromAccount.Balance < amount {
		return errors.New("insufficient balance")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	addTransactionToPassbook(fromAccount, "transfer-debit", amount, fromAccount.Balance, toAccount.AccountID)
	addTransactionToPassbook(toAccount, "transfer-credit", amount, toAccount.Balance, fromAccount.AccountID)

	return nil
}

func addTransactionToPassbook(account *models.Account, transactionType string, amount, newBalance float64, correspondingAccount int) {
	transaction := models.Transaction{
		TransactionType:      transactionType,
		Amount:               amount,
		Time:                 time.Now(),
		NewBalance:           newBalance,
		CorrespondingAccount: correspondingAccount,
	}
	account.Passbook = append(account.Passbook, transaction)
}
