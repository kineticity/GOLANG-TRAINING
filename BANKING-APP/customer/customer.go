package customer

import (
	"bankingApp/account"
	"bankingApp/bank"
	"errors"
	"fmt"
	"bankingApp/validations"
)

var customerid int = 1
var users []*User

type User struct {
	customerID     int
	firstName      string
	lastName       string
	isAdmin        bool
	isBankCustomer bool
	isActive       bool
	accounts       []*account.Account
	totalBalance   float64
}

type Admin interface {
	NewCustomer(firstName, lastName string) (*User, error)
	GetCustomers() ([]*User, error)
	GetCustomerByID(customerID int) (*User, error)
	UpdateCustomer(customerID int, attribute string, newValue interface{}) error
	DeleteCustomer(customerID int) error
	CreateBank(fullName, abbreviation string) error
	GetAllBanks() ([]bank.BankOperations, error)
	GetBankByID(bankID int) (bank.BankOperations, error)
	UpdateBank(bankID int, param string, newValue string) error
	DeleteBank(bankID int) error
	//getter setter of customer?

}

type BankCustomer interface {
	CreateAccount(initialBalance float64, bankid int) error
	GetAccountByID(accountID int) (*account.Account, error)
	GetAllAccounts() ([]*account.Account, error)
	UpdateAccount(accountID int, attribute string, newValue interface{}) error
	DeleteAccount(accountID int) error
	UpdateTotalBalance()
	Deposit(accountNo int, amount float64) error
	Withdraw(accountNo int, amount float64) error
	Transfer(fromAccountNo, toAccountNo int, amount float64) error
}

func NewAdmin(firstName, lastName string) (*User, error) {

	if firstName == "" || lastName == "" {
		return nil, errors.New("first and last name cannot be empty")
	}

	c := &User{
		customerID:     customerid,
		firstName:      firstName,
		lastName:       lastName,
		isAdmin:        true,
		isBankCustomer: false,
		isActive:       true,
	}
	customerid++
	// customers=append(customers, c)

	return c, nil
}

func (c *User) NewCustomer(firstName, lastName string) (*User, error) { //CREATE CUSTOMER
	if !c.GetIsAdmin() {
		return nil, errors.New("only admin can create new customers")
	}


	if err := validation.ValidateNonEmptyString("First name", firstName); err != nil {
		return nil, err
	}
	if err := validation.ValidateNonEmptyString("Last name", lastName); err != nil {
		return nil, err
	}

	newCustomer := &User{
		customerID:     customerid,
		firstName:      firstName,
		lastName:       lastName,
		isAdmin:        false,
		isBankCustomer: true,
		isActive:       true,
	}
	customerid++
	users = append(users, newCustomer) //users is storing customers only not admins

	return newCustomer, nil
}

// Getter Setter fns
func (c *User) GetCustomerID() int {
	return c.customerID
}

func (c *User) SetCustomerID(id int) {
	c.customerID = id
}

func (c *User) GetFirstName() string {
	return c.firstName
}

func (c *User) SetFirstName(firstName string) {
	c.firstName = firstName
}

func (c *User) GetLastName() string {
	return c.lastName
}

func (c *User) SetLastName(lastName string) {
	c.lastName = lastName
}

func (c *User) GetIsAdmin() bool {
	return c.isAdmin
}

func (c *User) SetIsAdmin(admin bool) {
	c.isAdmin = admin
}
func (c *User) GetIsBankCustomer() bool {
	return c.isBankCustomer
}

func (c *User) SetIsBankCustomer(bankcustomer bool) {
	c.isBankCustomer = bankcustomer
}
func (c *User) GetIsActive() bool {
	return c.isActive
}

func (c *User) SetIsActive(active bool) {
	c.isActive = active
}

func (c *User) GetAccounts() []*account.Account {
	return c.accounts
}

func (c *User) SetAccounts(accounts []*account.Account) {
	c.accounts = accounts
}

func (c *User) GetTotalBalance() float64 {
	return c.totalBalance
}

func (c *User) SetTotalBalance(balance float64) {
	c.totalBalance = balance
}

func (c *User) GetCustomers() ([]*User, error) {
	if !c.GetIsAdmin() {
		return nil, errors.New("only admins can retrieve customers")
	}
	return users, nil
}

func (c *User) GetCustomerByID(customerID int) (*User, error) {
	if !c.GetIsAdmin() {
		return nil, errors.New("only admins can retrieve customers by ID")
	}

	for _, customer := range users {
		if customer.GetCustomerID() == customerID && customer.GetIsActive() {
			return customer, nil
		}
	}
	return nil, errors.New("customer not found or inactive")
}
func (c *User) UpdateCustomer(customerID int, attribute string, newValue interface{}) error { //UPDATE CUSTOMER
	if !c.GetIsAdmin() {
		return errors.New("only admins can update non-admin customers")
	}

	var targetCustomer *User
	for _, customer := range users {
		if customer.GetCustomerID() == customerID && !customer.GetIsAdmin() && customer.GetIsActive() { //admin cannot update other admins !customer.GetIsAdmin()
			targetCustomer = customer
			break
		}
	}
	if targetCustomer == nil {
		return errors.New("target customer not found or inactive")
	}

	switch attribute {
	case "firstName":

		if value, ok := newValue.(string); ok {
			if err := validation.ValidateNonEmptyString("First name", value); err != nil {
				return err
			}
			targetCustomer.SetFirstName(value)
			return nil
		}
		return errors.New("invalid value for first name")
	case "lastName":

		if value, ok := newValue.(string); ok {
			if err := validation.ValidateNonEmptyString("Last name", value); err != nil {
				return err
			}
			targetCustomer.SetLastName(value)
			return nil
		}
		return errors.New("invalid value for last name")
		
	default:
		return errors.New("unknown attribute")
	}
}

func (c *User) DeleteCustomer(customerID int) error { //DELETE CUSTOMER
	if !c.GetIsAdmin() {
		return errors.New("only admins can delete non-admin customers")
	}

	for _, customer := range users {
		if customer.GetCustomerID() == customerID && !customer.GetIsAdmin() && customer.GetIsActive() {
			customer.SetIsActive(false)
			fmt.Printf("Customer %s %s has been deleted.\n", customer.GetFirstName(), customer.GetLastName())
			return nil
		}
	}
	return errors.New("non-admin customer not found or already inactive")
}

func (c *User) CreateAccount(initialBalance float64, bankid int) error { //CREATE ACCOUNT
	if !c.GetIsBankCustomer() {
		return errors.New("only bank customers can create accounts")
	}
	if !c.GetIsActive() {
		return errors.New("customer is not active, cannot create account")
	}
	bankWhereAccountIsMade, err := bank.GetBankByID(bankid) //should this be in account?
	if err != nil {
		return err
	}

	newAccount, err := account.NewAccount(initialBalance, bankid)
	if err != nil {
		return err
	}

	bankWhereAccountIsMade.AddAccount(newAccount)
	c.SetAccounts(append(c.GetAccounts(), newAccount))
	c.UpdateTotalBalance()
	return nil
}

func (c *User) GetAccountByID(accountID int) (*account.Account, error) { //GETBYID
	if !c.GetIsBankCustomer() {
		return nil, errors.New("only bank customers can retrieve their accounts")
	}
	if !c.GetIsActive() {
		return nil, errors.New("customer is not active, cannot retrieve account")
	}

	return account.GetAccountByID(accountID, c.GetAccounts())
}

func (c *User) GetAllAccounts() ([]*account.Account, error) { //GETALL
	if !c.GetIsBankCustomer() {
		return nil, errors.New("only bank customers can retrieve their accounts")
	}
	if !c.GetIsActive() {
		return nil, errors.New("customer is not active, cannot retrieve accounts")
	}

	return c.GetAccounts(), nil
}

func (c *User) UpdateAccount(accountID int, attribute string, newValue interface{}) error { //UPDATE ACCOUNT
	if !c.GetIsBankCustomer() {
		return errors.New("only bank customers can update their accounts")
	}
	if !c.GetIsActive() {
		return errors.New("customer is not active, cannot update account")
	}

	var acc *account.Account
	for _, account := range c.GetAccounts() {
		if account.GetAccountID() == accountID {
			acc = account
			break
		}
	}
	if acc == nil || !acc.GetIsActive() {
		return errors.New("account not found or inactive")
	}

	return acc.UpdateAccount(attribute, newValue)
}

func (c *User) DeleteAccount(accountID int) error { //DELETE ACCOUNT
	if !c.GetIsBankCustomer() {
		return errors.New("only bank customers can delete their accounts")
	}
	if !c.GetIsActive() {
		return errors.New("customer is not active, cannot delete account")
	}

	for _, acc := range c.GetAccounts() {
		if acc.GetAccountID() == accountID {
			acc.Deactivate()
			c.UpdateTotalBalance()
			fmt.Printf("Account ID %d has been deleted.\n", accountID)
			return nil
		}
	}
	return errors.New("account not found")
}

func (c *User) UpdateTotalBalance() {
	total := 0.0
	for _, acc := range c.GetAccounts() {
		if acc.GetIsActive() {
			total += acc.GetBalance()
		}
	}
	c.SetTotalBalance(total)
}

func (c *User) CreateBank(fullName, abbreviation string) error { //CREATE BANK
	if !c.GetIsAdmin() {
		return errors.New("only admin can create banks")
	}

	_, err := bank.NewBank(fullName, abbreviation)
	if err != nil {
		return err
	}

	return nil
}

func (c *User) GetAllBanks() ([]bank.BankOperations, error) { //GETALL
	if !c.GetIsAdmin() {
		return nil, errors.New("only admin can read all banks")
	}
	// var banks bank.BankOperations
	banks := bank.GetAllBanks()

	if len(banks) == 0 {
		return nil, errors.New("no banks available")
	}

	return banks, nil
}


func (c *User) GetBankByID(bankID int) (bank.BankOperations, error) { //GETBYID
	if !c.GetIsAdmin() {
		return nil, errors.New("only admin can read a bank by ID")
	}
	var b bank.BankOperations
	b, err := bank.GetBankByID(bankID)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *User) UpdateBank(bankID int, param string, newValue string) error { //UPDATE BANK
	if !c.GetIsAdmin() {
		return errors.New("only admin can update banks")
	}
	var bankToUpdate bank.BankOperations //DIP
	bankToUpdate, err := bank.GetBankByID(bankID)
	if err != nil {
		return err
	}

	err = bankToUpdate.UpdateBankField(param, newValue)
	if err != nil {
		return fmt.Errorf("error updating bank: %v", err)
	}

	fmt.Printf("Admin updated bank %d's %s to %s\n", bankID, param, newValue)
	return nil
}

// DELETE BANK BY ID
func (c *User) DeleteBank(bankID int) error {
	if !c.GetIsAdmin() {
		return errors.New("only admin can delete banks")
	}
	var bankToDelete bank.BankOperations
	bankToDelete,err:=bank.GetBankByID(bankID)
	if err!=nil{
		return err
	}
	bankToDelete.Delete()

	return nil
}


func (c *User) Deposit(accountNo int, amount float64) error {
	if !c.GetIsBankCustomer() {
		return errors.New("only bank customers can deposit money")
	}
	if !c.GetIsActive() {
		return errors.New("inactive customer cannot perform transactions")
	}

	acc, err := account.GetAccountByID(accountNo, c.GetAccounts())
	if err != nil {
		return err
	}
	if !acc.GetIsActive() {
		return errors.New("cannot deposit to an inactive account")
	}
	bankid := acc.GetBankID()

	err = acc.Deposit(amount, accountNo, bankid, false, accountNo, bankid)
	if err != nil {
		return err
	}

	fmt.Printf("Deposited Rs. %.2f to account %d. New Balance: %.2f\n", amount, accountNo, acc.GetBalance())
	return nil
}

func (c *User) Withdraw(accountNo int, amount float64) error {
	if !c.GetIsBankCustomer() {
		return errors.New("only bank customers can withdraw money")
	}
	if !c.GetIsActive() {
		return errors.New("inactive customer cannot perform transactions")
	}

	acc, err := account.GetAccountByID(accountNo, c.GetAccounts())
	if err != nil {
		return err
	}
	if !acc.GetIsActive() {
		return errors.New("cannot withdraw from an inactive account")
	}
	bankid := acc.GetBankID()

	err = acc.Withdraw(amount, accountNo, bankid, false, accountNo, bankid)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrew Rs. %.2f from account %d. New Balance: %.2f\n", amount, accountNo, acc.GetBalance())
	return nil
}

func (c *User) Transfer(fromAccountNo, toAccountNo int, amount float64) error {
	if !c.GetIsBankCustomer() {
		return errors.New("only bank customers can transfer money between accounts")
	}
	if !c.GetIsActive() {
		return errors.New("customer is inactive, cannot perform transfers")
	}

	fromAcc, err := account.GetAccountByID(fromAccountNo, c.GetAccounts())
	if err != nil {
		return err
	}

	fromAccBankid := fromAcc.GetBankID()

	toAcc, err := account.GetAccountByID(toAccountNo, c.GetAccounts())
	if err == nil {
		if !toAcc.GetIsActive() {
			return errors.New("destination account is inactive")
		}
	} else {
		toAcc, err = account.GetAccountByID(toAccountNo, account.GetAllAccounts()) //global allaccounts not just customers accounts
		if err != nil {
			return errors.New("destination account not found or inactive")
		}
	}

	toAccBankid := toAcc.GetBankID()

	if !fromAcc.GetIsActive() || !toAcc.GetIsActive() {
		return errors.New("both accounts must be active")
	}

	err = fromAcc.Withdraw(amount, fromAccountNo, fromAccBankid, true, toAccountNo, toAccBankid)
	if err != nil {
		return err
	}

	err = toAcc.Deposit(amount, toAccountNo, toAccBankid, true, fromAccountNo, fromAccBankid)
	if err != nil {
		return err
	}

	fmt.Printf("Transferred Rs. %.2f from account %d to account %d\n", amount, fromAccountNo, toAccountNo)
	return nil
}
