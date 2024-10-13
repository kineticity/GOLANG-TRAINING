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
	customerID   int
	firstName    string
	lastName     string
	isAdmin      bool
	isActive     bool
	accounts     []*account.Account
	totalBalance float64
}


func NewAdmin(firstName, lastName string) (*Customer, error) {



	if firstName == "" || lastName == "" {
		return nil, errors.New("first and last name cannot be empty")
	}
	c:=&Customer{
		customerID: customerid,
		firstName:  firstName,
		lastName:   lastName,
		isAdmin:    true,
		isActive:   true, 
	}
	customerid++
	// customers=append(customers, c)


	return c, nil
}

func (c *Customer) NewCustomer(firstName, lastName string) (*Customer, error) { //CREATE CUSTOMER
	if !c.GetIsAdmin() {
		return nil, errors.New("only admin can create new customers")
	}



	if firstName == "" || lastName == "" {
		return nil, errors.New("first and last name cannot be empty")
	}

	newCustomer := &Customer{
		customerID: customerid,
		firstName:  firstName,
		lastName:   lastName,
		isAdmin:    false, 
		isActive:   true,   
	}
	customerid++
	customers=append(customers, newCustomer)

	return newCustomer, nil
}

// Getter Setter fns
func (c *Customer) GetCustomerID() int {
	return c.customerID
}

func (c *Customer) SetCustomerID(id int) {
	c.customerID = id
}

func (c *Customer) GetFirstName() string {
	return c.firstName
}

func (c *Customer) SetFirstName(firstName string) {
	c.firstName = firstName
}

func (c *Customer) GetLastName() string {
	return c.lastName
}

func (c *Customer) SetLastName(lastName string) {
	c.lastName = lastName
}

func (c *Customer) GetIsAdmin() bool {
	return c.isAdmin
}

func (c *Customer) SetIsAdmin(admin bool) {
	c.isAdmin = admin
}

func (c *Customer) GetIsActive() bool {
	return c.isActive
}

func (c *Customer) SetIsActive(active bool) {
	c.isActive = active
}

func (c *Customer) GetAccounts() []*account.Account {
	return c.accounts
}

func (c *Customer) SetAccounts(accounts []*account.Account) {
	c.accounts = accounts
}

func (c *Customer) GetTotalBalance() float64 {
	return c.totalBalance
}

func (c *Customer) SetTotalBalance(balance float64) {
	c.totalBalance = balance
}

func (c *Customer) GetCustomers() ([]*Customer, error) {
	if !c.GetIsAdmin() {
		return nil, errors.New("only admins can retrieve customers")
	}
	return customers,nil
}

func (c *Customer) GetCustomerByID(customerID int) (*Customer, error) {
	if !c.GetIsAdmin() {
		return nil, errors.New("only admins can retrieve customers by ID")
	}

	for _, customer := range customers {
		if customer.GetCustomerID() == customerID && customer.GetIsActive() {
			return customer, nil
		}
	}
	return nil, errors.New("customer not found or inactive")
}
func (c *Customer) UpdateCustomer(customerID int, attribute string, newValue interface{}) error { //UPDATE CUSTOMER
	if !c.GetIsAdmin() {
		return errors.New("only admins can update non-admin customers")
	}

	var targetCustomer *Customer
	for _, customer := range customers {
		if customer.GetCustomerID() == customerID && !customer.GetIsAdmin() && customer.GetIsActive() { //admin cannot update other admins !customer.GetIsAdmin()
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
			targetCustomer.SetFirstName(value)
			return nil
		}
		return errors.New("invalid value for first name")
	case "lastName":
		if value, ok := newValue.(string); ok && value != "" {
			targetCustomer.SetLastName(value) 
			return nil
		}
		return errors.New("invalid value for last name")
	default:
		return errors.New("unknown attribute")
	}
}

func (c *Customer) DeleteCustomer(customerID int) error { //DELETE CUSTOMER
	if !c.GetIsAdmin() {
		return errors.New("only admins can delete non-admin customers")
	}

	for _, customer := range customers {
		if customer.GetCustomerID() == customerID && !customer.GetIsAdmin() && customer.GetIsActive() {
			customer.SetIsActive(false)
			fmt.Printf("Customer %s %s has been deleted.\n", customer.GetFirstName(), customer.GetLastName())
			return nil
		}
	}
	return errors.New("non-admin customer not found or already inactive")
}

func (c *Customer) CreateAccount(initialBalance float64,bankid int) error { //CREATE ACCOUNT
	if c.GetIsAdmin() {
		return errors.New("admins cannot create accounts")
	}
	if !c.GetIsActive() {
		return errors.New("customer is not active, cannot create account")
	}
	if initialBalance < 1000 {
		return errors.New("initial balance must be at least Rs. 1000")
	}

	newAccount, err := account.NewAccount(initialBalance,bankid)
	if err != nil {
		return err
	}

	bankWhereAccountIsMade,err:=bank.GetBankByID(bankid)
	if err!=nil{
		return err
	}
	bankWhereAccountIsMade.AddAccount(newAccount)
	c.SetAccounts(append(c.GetAccounts(),newAccount))
	c.UpdateTotalBalance() 
	return nil
}

func (c *Customer) GetAccountByID(accountID int) (*account.Account, error) { //GETBYID
	if c.GetIsAdmin() {
		return nil, errors.New("admins cannot retrieve accounts")
	}
	if !c.GetIsActive() {
		return nil, errors.New("customer is not active, cannot retrieve account")
	}

	return account.GetAccountByID(accountID, c.GetAccounts())
}

func (c *Customer) GetAllAccounts() ([]*account.Account, error) { //GETALL
	if c.GetIsAdmin() {
		return nil, errors.New("admins cannot retrieve accounts")
	}
	if !c.GetIsActive() {
		return nil, errors.New("customer is not active, cannot retrieve accounts")
	}

	return c.GetAccounts(), nil
}

func (c *Customer) UpdateAccount(accountID int, attribute string, newValue interface{}) error { //UPDATE ACCOUNT
	if c.GetIsAdmin() {
		return errors.New("admins cannot update accounts")
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

func (c *Customer) DeleteAccount(accountID int) error { //DELETE ACCOUNT
	if c.GetIsAdmin() {
		return errors.New("admins cannot delete accounts")
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

func (c *Customer) UpdateTotalBalance() {
	total := 0.0
	for _, acc := range c.GetAccounts() {
		if acc.GetIsActive() {
			total += acc.GetBalance()
		}
	}
	c.SetTotalBalance(total)
}

func (c *Customer) CreateBank(bankID int, fullName, abbreviation string) error { //CREATE BANK
	if !c.GetIsAdmin() {
		return errors.New("only admin can create banks")
	}

	_, err := bank.NewBank(fullName, abbreviation)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) GetAllBanks() ([]*bank.Bank, error) { //GETALL
	if !c.GetIsAdmin() {
		return nil, errors.New("only admin can read all banks")
	}

	banks := bank.GetAllBanks()

	if len(banks) == 0 {
		return nil, errors.New("no banks available")
	}

	return banks, nil
}

func (c *Customer) GetBankByID(bankID int) (*bank.Bank, error) { //GETBYID
	if !c.GetIsAdmin() {
		return nil, errors.New("only admin can read a bank by ID")
	}

	b, err := bank.GetBankByID(bankID)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Customer) UpdateBank(bankID int, param string, newValue string) error { //UPDATE BANK
	if !c.GetIsAdmin() {
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
	if !c.GetIsAdmin() {
		return errors.New("only admin can delete banks")
	}

	err := bank.DeleteBankByID(bankID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Customer) Deposit(accountNo int, amount float64) error {
	if c.GetIsAdmin() {
		return errors.New("admins cannot deposit money")
	}
	if !c.GetIsActive() {
		return errors.New("inactive customer cannot perform transactions")
	}

	acc, err := account.GetAccountByID(accountNo,c.GetAccounts())
	if err != nil {
		return err
	}
	if !acc.GetIsActive() {
		return errors.New("cannot deposit to an inactive account")
	}
	bankid:=acc.GetBankID()

	err = acc.Deposit(amount,accountNo,bankid,false,accountNo,bankid)
	if err != nil {
		return err
	}

	fmt.Printf("Deposited Rs. %.2f to account %d. New Balance: %.2f\n", amount, accountNo, acc.GetBalance())
	return nil
}

func (c *Customer) Withdraw(accountNo int, amount float64) error {
	if c.GetIsAdmin() {
		return errors.New("admins cannot withdraw money")
	}
	if !c.GetIsActive() {
		return errors.New("inactive customer cannot perform transactions")
	}

	acc, err := account.GetAccountByID(accountNo,c.GetAccounts())
	if err != nil {
		return err
	}
	if !acc.GetIsActive() {
		return errors.New("cannot withdraw from an inactive account")
	}
	bankid:=acc.GetBankID()


	err = acc.Withdraw(amount,accountNo,bankid,false,accountNo,bankid)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrew Rs. %.2f from account %d. New Balance: %.2f\n", amount, accountNo, acc.GetBalance())
	return nil
}

func (c *Customer) Transfer(fromAccountNo, toAccountNo int, amount float64) error {
	if c.GetIsAdmin() {
		return errors.New("admins cannot transfer money between accounts")
	}
	if !c.GetIsActive() {
		return errors.New("customer is inactive, cannot perform transfers")
	}

	fromAcc, err := account.GetAccountByID(fromAccountNo,c.GetAccounts())
	if err != nil {
		return err
	}

	fromAccBankid:=fromAcc.GetBankID()


	toAcc, err := account.GetAccountByID(toAccountNo,c.GetAccounts())
	if err == nil {
		if !toAcc.GetIsActive() {
			return errors.New("destination account is inactive")
		}
	} else {
		toAcc, err = account.GetAccountByID(toAccountNo,account.GetAllAccounts()) //global allaccounts not just customers accounts
		if err != nil {
			return errors.New("destination account not found or inactive")
		}
	}

	toAccBankid:=toAcc.GetBankID()

	if !fromAcc.GetIsActive() || !toAcc.GetIsActive() {
		return errors.New("both accounts must be active")
	}

	err = fromAcc.Withdraw(amount,fromAccountNo,fromAccBankid,true,toAccountNo,toAccBankid)
	if err != nil {
		return err
	}

	err = toAcc.Deposit(amount,toAccountNo,toAccBankid,true,fromAccountNo,fromAccBankid)
	if err != nil {
		return err
	}

	fmt.Printf("Transferred Rs. %.2f from account %d to account %d\n", amount, fromAccountNo, toAccountNo)
	return nil
}