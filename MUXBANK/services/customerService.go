package services

import (
	// "bankingApp/controllers"
	"bankingApp/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)


var allCustomers = []*models.User{{UserID: 2, Username: "customer", Password: string(models.Cutomerpasshash), IsAdmin: false, IsCustomer: true}}
var customerIDCounter = 3

func CreateCustomer(username, password, firstName, lastName string) (*models.User, error) {
	for _, customer := range allCustomers {
		if customer.Username == username {
			return nil, errors.New("username already exists")
		}
	}
	hashedpass,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return nil,err
	}
	customer := &models.User{
		UserID:     customerIDCounter,
		Username:   username,
		Password:   string(hashedpass),
		FirstName:  firstName,
		LastName:   lastName,
		IsAdmin:    false,
		IsCustomer: true,
		Accounts:   []*models.Account{}, // Initialize accounts slice
	}

	customerIDCounter++
	allCustomers = append(allCustomers, customer)         //jsut customers
	models.Userslist = append(models.Userslist, customer) //admin+customers list
	return customer, nil
}

func GetCustomerByID(userID int) (*models.User, error) {
	for _, customer := range allCustomers {
		if customer.UserID == userID {
			return customer, nil
		}
	}
	return nil, errors.New("customer not found")
}

func GetAllCustomers() []*models.User {
	return allCustomers
}

func UpdateCustomer(userID int, username, password, firstName, lastName string) (*models.User, error) {
	for i, customer := range allCustomers {
		if customer.UserID == userID {
			allCustomers[i].Username = username
			allCustomers[i].Password = password
			allCustomers[i].FirstName = firstName
			allCustomers[i].LastName = lastName
			return allCustomers[i], nil
		}
	}
	return nil, errors.New("customer not found")
}

func DeleteCustomer(userID int) error {
	for i, customer := range allCustomers {
		if customer.UserID == userID {
			allCustomers = append(allCustomers[:i], allCustomers[i+1:]...)
			return nil
		}
	}
	return errors.New("customer not found")
}
