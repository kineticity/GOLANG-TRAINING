package main

import (
	"fmt"
	// "bankingApp/account"
	"bankingApp/bank"
	"bankingApp/customer"
)

func main() {
	// Create an admin customer
	admin, err := customer.NewAdmin("Sean", "Smith")
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	fmt.Println("Admin created:", admin.GetFirstName(), admin.GetLastName())

	// Create a new bank
	bank1, err := bank.NewBank("National Bank", "NB")
	if err != nil {
		fmt.Println("Error creating bank:", err)
		return
	}
	fmt.Println("Bank created:", bank1.GetFullName())

	bank2, err := bank.NewBank("American Bank", "AB")
	if err != nil {
		fmt.Println("Error creating bank:", err)
		return
	}
	fmt.Println("Bank created:", bank2.GetFullName())

	// Create a new customer by admin
	newCustomer, err := admin.NewCustomer("John", "Doe")
	if err != nil {
		fmt.Println("Error creating customer:", err)
		return
	}
	fmt.Println("New customer created:", newCustomer.GetFirstName(), newCustomer.GetLastName())

	// Create an account for the new customer
	err = newCustomer.CreateAccount(1500, bank1.GetBankID()) //Account 1 of newCustomer
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	fmt.Println("Account created for customer:", newCustomer.GetFirstName())

	err = newCustomer.CreateAccount(2000, bank1.GetBankID()) //Account 2 of newCustomer
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	fmt.Println("Account created for customer:", newCustomer.GetFirstName())

	err = newCustomer.CreateAccount(2000, bank2.GetBankID()) //Account 3 of newCustomer in different bank
	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}
	fmt.Println("Account created for customer:", newCustomer.GetFirstName())

	// Retrieve the account by ID
	accountID := newCustomer.GetAccounts()[0].GetAccountID()
	acc, err := newCustomer.GetAccountByID(accountID)
	if err != nil {
		fmt.Println("Error retrieving account:", err)
		return
	}
	fmt.Printf("Retrieved Account ID: %d, Balance: %.2f\n", acc.GetAccountID(), acc.GetBalance())

	// Deposit money into the account
	err = newCustomer.Deposit(accountID, 500)
	if err != nil {
		fmt.Println("Error depositing:", err)
		return
	}

	// Withdraw money from the account
	err = newCustomer.Withdraw(accountID, 200)
	if err != nil {
		fmt.Println("Error withdrawing:", err)
		return
	}

	// Transfer money between accounts
	if len(newCustomer.GetAccounts()) > 1 { //self transfer
		anotherAccountID := newCustomer.GetAccounts()[1].GetAccountID()
		fmt.Println(anotherAccountID)
		err = newCustomer.Transfer(accountID, anotherAccountID, 100)
		if err != nil {
			fmt.Println("Error transferring:", err)
			return
		}

		anotherAccountIDInAnotherBank := newCustomer.GetAccounts()[2].GetAccountID()
		anotherBankIDInAnotherBank := newCustomer.GetAccounts()[2].GetBankID()

		// fmt.Println(anotherAccountIDInAnotherBank)
		fmt.Println(anotherBankIDInAnotherBank)

		err = newCustomer.Transfer(accountID, anotherAccountIDInAnotherBank, 300)
		if err != nil {
			fmt.Println("Error transferring:", err)
			return
		}

		fmt.Println(anotherBankIDInAnotherBank)

	}

	passbook := acc.GetPassbook() // Assuming there's a method to get the passbook
	passbook.PrintPassbook()

	// List all customers (admin functionality)
	customers, err := admin.GetCustomers()
	if err != nil {
		fmt.Println("Error retrieving customers:", err)
		return
	}
	fmt.Println("List of customers:")
	for _, c := range customers {
		fmt.Printf("%s %s\n", c.GetFirstName(), c.GetLastName())
	}

	// Delete the customer
	err = admin.DeleteCustomer(newCustomer.GetCustomerID())
	if err != nil {
		fmt.Println("Error deleting customer:", err)
		return
	}
	fmt.Println("Customer deleted:", newCustomer.GetFirstName(), newCustomer.GetLastName(), newCustomer.GetIsActive())

	sbi, _ := bank.NewBank("State Bank of India", "SBI")
	icici, _ := bank.NewBank("ICICI Bank", "ICICI")

	sbi.LendTo("ICICI", 1000, 4)
	sbi.ReceiveFrom("ICICI", 500, 4)

	fmt.Println("SBI Ledger:")
	sbi.PrintBankLedger()

	fmt.Println("\nICICI Ledger:")
	icici.PrintBankLedger()

	var admindemo customer.Admin



	admindemo,err=customer.NewAdmin("Admin","2")
	if err!=nil{
		fmt.Println(err)
	}
	admindemo.CreateBank("Punjab National Bank","PNB")
	// banks,_:=admindemo.GetAllBanks()
	// for _,bank:= range banks{
	// 	fmt.Println(bank.GetBankID())
	// 	fmt.Println(bank.GetFullName())
	// 	fmt.Println(bank.GetAbbreviation())
	// 	fmt.Println(bank.GetIsActive())
	// 	fmt.Println(bank.GetLedger())
	// 	bank.GetLedger().PrintLedger()	
	// }
	bank,err:=admindemo.GetBankByID(5)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(bank.Read())
	}
	err=admindemo.DeleteBank(5)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Deleted successfully")
	}

	// fmt.Println(err)
	bank,err=admindemo.GetBankByID(5)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(bank.Read())
	}

	bankopobj,_:=admindemo.GetAllBanks()
	fmt.Println(bankopobj[0].GetFullName())
	// fmt.Println(bank.Read())
}

