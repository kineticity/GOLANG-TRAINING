package models

import ("time"
"golang.org/x/crypto/bcrypt")


var adminpasshash,_=bcrypt.GenerateFromPassword([]byte("adminpass"),bcrypt.DefaultCost)
var Cutomerpasshash,_=bcrypt.GenerateFromPassword([]byte("customerpass"),bcrypt.DefaultCost)

var Userslist = []*User{
	{UserID: 1, Username: "admin", Password: string(adminpasshash), IsAdmin: true, IsCustomer: false},
	{UserID: 2, Username: "customer", Password: string(Cutomerpasshash), IsAdmin: false, IsCustomer: true},
}


type Ledger struct{
	LedgerEntries []*LedgerEntry `json:"ledgerentry"`
}


type User struct {
	UserID     int        `json:"userID"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	IsAdmin    bool       `json:"isAdmin"`
	IsCustomer bool       `json:"isCustomer"`
	Accounts   []*Account `json:"accounts"` 
}



type LedgerEntry struct {
	BankName string  `json:"bankname"`
	Amount   float64 `json:"amount"`
	Time     time.Time `json:"time"`
}

type Bank struct {
	BankID      int           `json:"bankid"`
	FullName    string        `json:"fullname"`
	Abbreviation string       `json:"abbreviation"`
	IsActive    bool          `json:"isactive"`
	Ledger      []LedgerEntry `json:"ledger"` 
}

type Transaction struct {
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Time            time.Time `json:"time"`
	NewBalance      float64   `json:"new_balance"`
	CorrespondingAccount int `json:"corresponding_account"` 
}

type Account struct {
	AccountID int         `json:"accountid"`
	CustomerID int        `json:"customerid"`
	BankID     int        `json:"bankid"`
	Balance    float64    `json:"balance"`
	Passbook   []Transaction `json:"passbook"` 
	IsActive bool
}
