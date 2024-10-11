package main

import (
	"fmt"
	"bankingApp/account"
	"bankingApp/customer"
)

func main() {
	// Create an Admin customer
	admin, err := customer.NewAdmin("Admin", "User")
	if err != nil {
		fmt.Println("Admin was not created:",err)
	}
	fmt.Println("Admin created successfully")

	// Admin creates a new bank
	err = admin.CreateBank(1, "National Bank", "NB")
	if err != nil {
		fmt.Println("Bank was not created:",err)
	}
	fmt.Println("Bank created successfully")

	//Admin creates new customers
	customer1, err := admin.NewCustomer("John", "Doe")
	if err != nil {
		fmt.Println("Customer was not created:",err)
	}
	admin.UpdateCustomer(2,"firstName","Jack")
	c,_:=admin.GetCustomerByID(2)
	fmt.Println(c.FirstName)
	// admin.DeleteCustomer(2)
	// _,err=admin.GetCustomerByID(2)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(c.FirstName)

	customer2, err := admin.NewCustomer("Jane", "Smith")
	if err != nil {
		fmt.Println("Error creating customer2:",err)
	}
	fmt.Println("Customer created successfully",customer2.FirstName)

	//Customer creates accounts
	err = customer1.CreateAccount(5000) //customer 1 account 1
	if err != nil {
	fmt.Println("Customer 1 Account 1 not created",customer1.FirstName)
	}
	err = customer1.CreateAccount(10000) //customer 1 account 2
	if err != nil {
	fmt.Println("Customer 1 account 2 not created",customer1.FirstName)
	}
	err = customer2.CreateAccount(7000) //customer 2 account 3
	if err != nil {
		fmt.Println("Customer 2 account 1 not created:",err)
	}

	//customer updates accounts
	err=customer1.UpdateAccount(1,"isActive",false)
	if err!=nil{
		fmt.Println(err)
	}else{
		a1,_:=customer1.GetAccountByID(1)
		fmt.Println(a1.IsActive)

	}


	// Customer deposits money
	err = customer1.Deposit(1, 2000)
	if err != nil {
		fmt.Println("Error in deposit for customer1:",err)
	}
	err = customer2.Deposit(3, 3000)
	if err != nil {
		fmt.Println("Error in deposit for customer2:",err)
	}
	fmt.Println("Deposit operation successful")

	//  Customer withdraws money
	err = customer1.Withdraw(1, 1000)
	if err != nil {
		fmt.Println("Error in withdrawal for customer1:",err)
	}
	err = customer2.Withdraw(3, 1500)
	if err != nil {
		fmt.Println("Error in withdrawal for customer2:",err)
	}
	fmt.Println("Withdrawal operation successful")

	// Transfer between accounts (self or another customer)
	err = customer1.Transfer(1, 2, 500) // Self-transfer within customer1's accounts acc 1 6000->5500 acc2 10000->10500 ,total customer1->16000
	if err != nil {
		fmt.Println("Error in self transfer for customer1:",err)
	}
	acc1,_:=account.GetGlobalAccountByID(1)
	fmt.Println("Account 1 balance is now:",acc1.Balance)
	acc2,_:=account.GetGlobalAccountByID(2)
	fmt.Println("Account 2 balance is now:",acc2.Balance)
	err = customer1.Transfer(1, 3, 1000) // Transfer to customer2's account acc1->5500->4500 acc3->8500->9500
	if err != nil {
		fmt.Println("Error in transfer from customer 1 to customer2:",err)
	}
	fmt.Println("Transfer operations successful")
	acc1,_=account.GetGlobalAccountByID(1)
	fmt.Println("Account 1 balance is now:",acc1.Balance)
	acc3,_:=account.GetGlobalAccountByID(3)
	fmt.Println("Account 3 balance is now:",acc3.Balance)

	fmt.Println("Customer 1 balance:",customer1.TotalBalance) //should be 15000

	// Admin updates bank details
	err = admin.UpdateBank(1, "bankName", "Updated National Bank")
	if err != nil {
		fmt.Println("Error updating bank details:",err)
	}
	fmt.Println("Bank details updated successfully")

	// Admin deletes customer2
	err = admin.DeleteCustomer(2)
	if err != nil {
		fmt.Println("Error deleting customer2:",err)
	}
	fmt.Println("Customer2 deleted successfully")

	// Retrieve and display all banks
	banks, err := admin.GetAllBanks()
	if err != nil {
		fmt.Println("Error retrieving banks:",err)
	}
	for _, b := range banks {
		fmt.Printf("Bank ID: %d, Name: %s, Abbreviation: %s, IsActive: %t\n", b.BankID, b.FullName, b.Abbreviation, b.IsActive)
	}

	// Retrieve and display all customers (non-admin)
	customers, err := admin.GetCustomers()
	if err != nil {
		fmt.Println("Error retrieving customers:",err)
	}
	for _, c := range customers {
		fmt.Printf("Customer ID: %d, Name: %s %s, Total Balance: %.2f, IsActive: %t\n", c.CustomerID, c.FirstName, c.LastName, c.TotalBalance, c.IsActive)
	}
}

