package services

import (
	database "bankingApp/databases"
	"bankingApp/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateCustomer(username, password, firstName, lastName string) (*models.User, error) {
	hashedpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	_, err = GetUserByUsername(username)
	if err == nil {
		return nil, fmt.Errorf("username already exists")
	}
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	user := &models.User{
		Username:   username,
		Password:   string(hashedpass),
		FirstName:  firstName,
		LastName:   lastName,
		IsAdmin:    false,
		IsCustomer: true,
		Accounts:   []*models.Account{},
	}

	err = user.Create(tx)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return nil, tx.Commit().Error
}

func DeleteCustomerByID(id int) error {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user := &models.User{Model: gorm.Model{ID: uint(id)}}
	if err := user.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func UpdateCustomerByID(id int, username, password, firstName, lastName string) (*models.User, error) {
	// Start a transaction
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	user := &models.User{}
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("customer not found: %v", err)
	}

	if username != "" && username != user.Username {
		_, err := GetUserByUsername(username)
		if err == nil {
			tx.Rollback()
			return nil, fmt.Errorf("username already exists")
		}
		user.Username = username
	}

	if password != "" {
		hashedpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("error hashing password: %v", err)
		}
		user.Password = string(hashedpass)
	}

	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}

	if err := user.Update(tx); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	return user, nil
}

func GetCustomerByID(id int) (*models.User, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := models.GetUserByID(tx, id)
	if err != nil {
		return nil, err
	}
	return user, nil

}
func GetAllCustomers() ([]models.User, error) {
	tx := database.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	return models.GetAllUsers(tx)
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := database.GetDB().Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
