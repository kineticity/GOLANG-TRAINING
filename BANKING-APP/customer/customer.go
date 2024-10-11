package customer

import (
	"errors"
	"fmt"
	"bankingApp/account"
	"bankingApp/bank"
)

var customerid int =1
var customers []*Customer

type Customer struct {
	CustomerID   int
	FirstName    string
	LastName     string
	IsAdmin      bool
	IsActive     bool
	Accounts     []*account.Account
	TotalBalance float64
}


func NewAdmin(firstName, lastName string) (*Customer, error) {



	if firstName == "" || lastName == "" {
		return nil, errors.New("first and last name cannot be empty")
	}
	c:=&Customer{
		CustomerID: customerid,
		FirstName:  firstName,
		LastName:   lastName,
		IsAdmin:    true,
		IsActive:   true, 
	}
	customerid++
	// customers=append(customers, c)


	return c, nil
}

func (c *Customer) NewCustomer(firstName, lastName string) (*Customer, error) { //CREATE CUSTOMER
	if !c.IsAdmin {
		return nil, errors.New("only admin can create new customers")
	}



	if firstName == "" || lastName == "" {
		return nil, errors.New("first and last name cannot be empty")
	}

	newCustomer := &Customer{
		CustomerID: customerid,
		FirstName:  firstName,
		LastName:   lastName,
		IsAdmin:    false, 
		IsActive:   true,   
	}
	customerid++
	customers=append(customers, newCustomer)

	return newCustomer, nil
}

func (c *Customer) GetCustomers() ([]*Customer, error) {
	if !c.IsAdmin {
		return nil, errors.New("only admins can retrieve customers")
	}

	var nonAdminCustomers []*Customer
	for _, customer := range customers {
		if !customer.IsAdmin && customer.IsActive {
			nonAdminCustomers = append(nonAdminCustomers, customer)
		} 
	}
	return nonAdminCustomers, nil
}

func (c *Customer) GetCustomerByID(customerID int) (*Customer, error) {
	if !c.IsAdmin {
		return nil, errors.New("only admins can retrieve customers by ID")
	}

	for _, customer := range customers {
		if customer.CustomerID == customerID && customer.IsActive {
			return customer, nil
		}
	}
	return nil, errors.New("customer not found or inactive")
}
func (c *Customer) UpdateCustomer(customerID int, attribute string, newValue interface{}) error { //UPDATE CUSTOMER
	if !c.IsAdmin {
		return errors.New("only admins can update non-admin customers")
	}

	var targetCustomer *Customer
	for _, customer := range customers {
		if customer.CustomerID == customerID && !customer.IsAdmin && customer.IsActive {
			targetCustomer = customer
			break
		}
	}
	if targetCustomer == nil {
		return errors.New("non-admin customer not found or inactive")
	}

	switch attribute {
	case "firstName":
		if value, ok := newValue.(string); ok && value != "" {
			targetCustomer.FirstName = value
			return nil
		}
		return errors.New("invalid value for first name")
	case "lastName":
		if value, ok := newValue.(string); ok && value != "" {
			targetCustomer.LastName = value
			return nil
		}
		return errors.New("invalid value for last name")
	default:
		return errors.New("unknown attribute")
	}
}

func (c *Customer) DeleteCustomer(customerID int) error { //DELETE CUSTOMER
	if !c.IsAdmin {
		return errors.New("only admins can delete non-admin customers")
	}

	for _, customer := range customers {
		if customer.CustomerID == customerID && !customer.IsAdmin && customer.IsActive {
			customer.IsActive = false // Mark as inactive
			fmt.Printf("Customer %s %s has been deleted.\n", customer.FirstName, customer.LastName)
			return nil
		}
	}
	return errors.New("non-admin customer not found or already inactive")
}

func (c *Customer) CreateAccount(initialBalance float64) error { //CREATE ACCOUNT
	if c.IsAdmin {
		return errors.New("admins cannot create accounts")
	}
	if !c.IsActive {
		return errors.New("customer is not active, cannot create account")
	}
	if initialBalance < 1000 {
		return errors.New("initial balance must be at least Rs. 1000")
	}

	newAccount, err := account.NewAccount(initialBalance)
	if err != nil {
		return err
	}

	c.Accounts = append(c.Accounts, newAccount)
	c.UpdateTotalBalance() 
	return nil
}

func (c *Customer) GetAccountByID(accountID int) (*account.Account, error) { //GETBYID
	if c.IsAdmin {
		return nil, errors.New("admins cannot retrieve accounts")
	}
	if !c.IsActive {
		return nil, errors.New("customer is not active, cannot retrieve account")
	}

	return account.GetAccountByID(accountID, c.Accounts)
}

func (c *Customer) GetAllAccounts() ([]*account.Account, error) { //GETALL
	if c.IsAdmin {
		return nil, errors.New("admins cannot retrieve accounts")
	}
	if !c.IsActive {
		return nil, errors.New("customer is not active, cannot retrieve accounts")
	}

	return c.Accounts, nil
}

func (c *Customer) UpdateAccount(accountID int, attribute string, newValue interface{}) error { //UPDATE ACCOUNT
	if c.IsAdmin {
		return errors.New("admins cannot update accounts")
	}
	if !c.IsActive {
		return errors.New("customer is not active, cannot update account")
	}

	var acc *account.Account
	for _, account := range c.Accounts {
		if account.AccountID == accountID {
			acc = account
			break
		}
	}
	if acc == nil || !acc.IsActive {
		return errors.New("account not found or inactive")
	}

	return acc.UpdateAccount(attribute, newValue)
}

func (c *Customer) DeleteAccount(accountID int) error { //DELETE ACCOUNT
	if c.IsAdmin {
		return errors.New("admins cannot delete accounts")
	}
	if !c.IsActive {
		return errors.New("customer is not active, cannot delete account")
	}

	for _, acc := range c.Accounts {
		if acc.AccountID == accountID {
			acc.Deactivate() 
			c.UpdateTotalBalance() 
			fmt.Printf("Account ID %d has been deleted.\n", accountID)
			return nil
		}
	}
	return errors.New("account not found")
}

func (c *Customer) UpdateTotalBalance() {
	total := 0.0
	for _, acc := range c.Accounts {
		if acc.IsActive {
			total += acc.Balance
		}
	}
	c.TotalBalance = total
}

func (c *Customer) CreateBank(bankID int, fullName, abbreviation string) error { //CREATE BANK
	if !c.IsAdmin {
		return errors.New("only admin can create banks")
	}

	_, err := bank.NewBank(fullName, abbreviation)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) GetAllBanks() ([]*bank.Bank, error) { //GETALL
	if !c.IsAdmin {
		return nil, errors.New("only admin can read all banks")
	}

	banks := bank.GetAllBanks()

	if len(banks) == 0 {
		return nil, errors.New("no banks available")
	}

	return banks, nil
}

func (c *Customer) GetBankByID(bankID int) (*bank.Bank, error) { //GETBYID
	if !c.IsAdmin {
		return nil, errors.New("only admin can read a bank by ID")
	}

	b, err := bank.GetBankByID(bankID)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Customer) UpdateBank(bankID int, param string, newValue string) error { //UPDATE BANK
	if !c.IsAdmin {
		return errors.New("only admin can update banks")
	}

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

//DELETE BANK BY ID
func (c *Customer) DeleteBank(bankID int) error {
	if !c.IsAdmin {
		return errors.New("only admin can delete banks")
	}

	err := bank.DeleteBankByID(bankID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) Deposit(accountNo int, amount float64) error {
	if c.IsAdmin {
		return errors.New("admins cannot deposit money")
	}
	if !c.IsActive {
		return errors.New("inactive customer cannot perform transactions")
	}

	acc, err := account.GetAccountByID(accountNo,c.Accounts)
	if err != nil {
		return err
	}
	if !acc.IsActive {
		return errors.New("cannot deposit to an inactive account")
	}

	err = acc.Deposit(amount)
	if err != nil {
		return err
	}

	fmt.Printf("Deposited Rs. %.2f to account %d. New Balance: %.2f\n", amount, accountNo, acc.Balance)
	return nil
}

func (c *Customer) Withdraw(accountNo int, amount float64) error {
	if c.IsAdmin {
		return errors.New("admins cannot withdraw money")
	}
	if !c.IsActive {
		return errors.New("inactive customer cannot perform transactions")
	}

	acc, err := account.GetAccountByID(accountNo,c.Accounts)
	if err != nil {
		return err
	}
	if !acc.IsActive {
		return errors.New("cannot withdraw from an inactive account")
	}

	err = acc.Withdraw(amount)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrew Rs. %.2f from account %d. New Balance: %.2f\n", amount, accountNo, acc.Balance)
	return nil
}

func (c *Customer) Transfer(fromAccountNo, toAccountNo int, amount float64) error {
	if c.IsAdmin {
		return errors.New("admins cannot transfer money between accounts")
	}
	if !c.IsActive {
		return errors.New("customer is inactive, cannot perform transfers")
	}

	fromAcc, err := account.GetAccountByID(fromAccountNo,c.Accounts)
	if err != nil {
		return err
	}

	toAcc, err := account.GetAccountByID(toAccountNo,c.Accounts)
	if err == nil {
		if !toAcc.IsActive {
			return errors.New("destination account is inactive")
		}
	} else {
		toAcc, err = account.GetGlobalAccountByID(toAccountNo)
		if err != nil {
			return errors.New("destination account not found or inactive")
		}
	}

	if !fromAcc.IsActive || !toAcc.IsActive {
		return errors.New("both accounts must be active")
	}

	err = fromAcc.Withdraw(amount)
	if err != nil {
		return err
	}

	err = toAcc.Deposit(amount)
	if err != nil {
		return err
	}

	fmt.Printf("Transferred Rs. %.2f from account %d to account %d\n", amount, fromAccountNo, toAccountNo)
	return nil
}