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

// DeleteNonAdminCustomer marks a non-admin customer as inactive
func (c *Customer) DeleteCustomer(customerID int) error { //DELETE CUSTOMER
	if !c.IsAdmin {
		return errors.New("only admins can delete non-admin customers")
	}

	// Find the non-admin customer by ID
	for _, customer := range customers {
		if customer.CustomerID == customerID && !customer.IsAdmin && customer.IsActive {
			customer.IsActive = false // Mark as inactive
			fmt.Printf("Customer %s %s has been deleted.\n", customer.FirstName, customer.LastName)
			return nil
		}
	}
	return errors.New("non-admin customer not found or already inactive")
}

// CreateAccount allows the customer to create a new account
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

	// // Determine the next account ID
	// accountID := 1
	// if len(c.Accounts) > 0 {
	// 	accountID = c.Accounts[len(c.Accounts)-1].AccountID + 1
	// }

	// Create a new account using the factory function
	newAccount, err := account.NewAccount(initialBalance)
	if err != nil {
		return err
	}

	// Append the new account to the customer's accounts
	c.Accounts = append(c.Accounts, newAccount)
	c.UpdateTotalBalance() // Update total balance after adding the new account
	return nil
}

// GetAccountByID retrieves an account by its ID
func (c *Customer) GetAccountByID(accountID int) (*account.Account, error) { //GETBYID
	if c.IsAdmin {
		return nil, errors.New("admins cannot retrieve accounts")
	}
	if !c.IsActive {
		return nil, errors.New("customer is not active, cannot retrieve account")
	}

	// Use the account package to get the account by ID
	return account.GetAccountByID(accountID, c.Accounts)
}

// GetAllAccounts retrieves all active accounts of the customer
func (c *Customer) GetAllAccounts() ([]*account.Account, error) { //GETALL
	if c.IsAdmin {
		return nil, errors.New("admins cannot retrieve accounts")
	}
	if !c.IsActive {
		return nil, errors.New("customer is not active, cannot retrieve accounts")
	}

	// Use the account package to get all accounts
	return c.Accounts, nil
}

// UpdateAccount updates the account's attributes
func (c *Customer) UpdateAccount(accountID int, attribute string, newValue interface{}) error { //UPDATE ACCOUNT
	if c.IsAdmin {
		return errors.New("admins cannot update accounts")
	}
	if !c.IsActive {
		return errors.New("customer is not active, cannot update account")
	}

	// Find the account by ID
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

	// Call the update method in the account package
	return acc.UpdateAccount(attribute, newValue)
}

// DeleteAccount marks an account as inactive
func (c *Customer) DeleteAccount(accountID int) error { //DELETE ACCOUNT
	if c.IsAdmin {
		return errors.New("admins cannot delete accounts")
	}
	if !c.IsActive {
		return errors.New("customer is not active, cannot delete account")
	}

	for _, acc := range c.Accounts {
		if acc.AccountID == accountID {
			acc.Deactivate() // Call the deactivate method from account package
			c.UpdateTotalBalance() // Update total balance after deletion
			fmt.Printf("Account ID %d has been deleted.\n", accountID)
			return nil
		}
	}
	return errors.New("account not found")
}

// UpdateTotalBalance updates the total balance of the customer
func (c *Customer) UpdateTotalBalance() {
	total := 0.0
	for _, acc := range c.Accounts {
		if acc.IsActive {
			total += acc.Balance
		}
	}
	c.TotalBalance = total
}

// CreateBank allows the admin to create a new bank and add it to the global banks slice
func (c *Customer) CreateBank(bankID int, fullName, abbreviation string) error { //CREATE BANK
	if !c.IsAdmin {
		return errors.New("only admin can create banks")
	}

	_, err := bank.NewBank(fullName, abbreviation)
	if err != nil {
		return err
	}

	// bank.AddBank(newBank)
	return nil
}

// ReadAllBanks allows the admin to read all banks
func (c *Customer) GetAllBanks() ([]*bank.Bank, error) { //GETALL
	if !c.IsAdmin {
		return nil, errors.New("only admin can read all banks")
	}

	// Fetch all banks from the bank package
	banks := bank.GetAllBanks()

	if len(banks) == 0 {
		return nil, errors.New("no banks available")
	}

	return banks, nil
}

// ReadBankByID allows the admin to read a bank by its ID
func (c *Customer) GetBankByID(bankID int) (*bank.Bank, error) { //GETBYID
	if !c.IsAdmin {
		return nil, errors.New("only admin can read a bank by ID")
	}

	// Fetch the bank by ID from the bank package
	b, err := bank.GetBankByID(bankID)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// UpdateBank allows the admin to update a specific field of the bank
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

// DeleteBank allows the admin to deactivate a bank from the global banks slice //DELETE BANK BY ID
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

// Deposit allows the customer to deposit money into an account
func (c *Customer) Deposit(accountNo int, amount float64) error {
	if c.IsAdmin {
		return errors.New("admins cannot deposit money")
	}
	if !c.IsActive {
		return errors.New("inactive customer cannot perform transactions")
	}

	// Get account by account ID
	acc, err := account.GetAccountByID(accountNo,c.Accounts)
	if err != nil {
		return err
	}
	if !acc.IsActive {
		return errors.New("cannot deposit to an inactive account")
	}

	// Perform deposit
	err = acc.Deposit(amount)
	if err != nil {
		return err
	}

	fmt.Printf("Deposited Rs. %.2f to account %d. New Balance: %.2f\n", amount, accountNo, acc.Balance)
	return nil
}

// Withdraw allows the customer to withdraw money from an account
func (c *Customer) Withdraw(accountNo int, amount float64) error {
	if c.IsAdmin {
		return errors.New("admins cannot withdraw money")
	}
	if !c.IsActive {
		return errors.New("inactive customer cannot perform transactions")
	}

	// Get account by account ID
	acc, err := account.GetAccountByID(accountNo,c.Accounts)
	if err != nil {
		return err
	}
	if !acc.IsActive {
		return errors.New("cannot withdraw from an inactive account")
	}

	// Perform withdraw
	err = acc.Withdraw(amount)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrew Rs. %.2f from account %d. New Balance: %.2f\n", amount, accountNo, acc.Balance)
	return nil
}

// Transfer transfers money between two accounts, either the customer's own accounts or between two different customers.
func (c *Customer) Transfer(fromAccountNo, toAccountNo int, amount float64) error {
	if c.IsAdmin {
		return errors.New("admins cannot transfer money between accounts")
	}
	if !c.IsActive {
		return errors.New("customer is inactive, cannot perform transfers")
	}

	// Get the 'from' account (must belong to the current customer)
	fromAcc, err := account.GetAccountByID(fromAccountNo,c.Accounts)
	if err != nil {
		return err
	}

	// Attempt to get the 'to' account from the current customer's accounts
	toAcc, err := account.GetAccountByID(toAccountNo,c.Accounts)
	if err == nil {
		// Self-transfer (if 'to' account is found in the current customer's accounts)
		if !toAcc.IsActive {
			return errors.New("destination account is inactive")
		}
	} else {
		// If the 'to' account is not found in the current customer's accounts, treat it as a transfer to another customer
		toAcc, err = account.GetGlobalAccountByID(toAccountNo)
		if err != nil {
			return errors.New("destination account not found or inactive")
		}
	}

	// Ensure both accounts are active
	if !fromAcc.IsActive || !toAcc.IsActive {
		return errors.New("both accounts must be active")
	}

	// Perform the transfer
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