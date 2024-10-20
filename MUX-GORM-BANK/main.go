package main

import (
	database "bankingApp/databases"
	"bankingApp/services"
	"fmt"

	// "bankingApp/models"
	"bankingApp/routes"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func main() {
	// var err error
	// DB, err := gorm.Open("mysql", "root:Forcepointpassword@1@/bankingappdb?charset=utf8&parseTime=True&loc=Local")
	// if err!=nil{
	// 	log.Fatalf("Could not connect to mysql db: %v", err)
	// }
	// defer DB.Close()

	// // Auto-migrate the models
	// DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Bank{}, &models.LedgerEntry{}, &models.Transaction{})
	// DB.Model(&models.Account{}).AddForeignKey("customer_id", "users(id)", "CASCADE", "SET NULL")
	// DB.Model(&models.Account{}).AddForeignKey("bank_id", "banks(id)", "CASCADE", "SET NULL")
	// DB.Model(&models.Transaction{}).AddForeignKey("account_id", "accounts(id)", "CASCADE", "SET NULL")

	hashedpass, _ := bcrypt.GenerateFromPassword([]byte("adminpass"), bcrypt.DefaultCost)
	fmt.Println(string(hashedpass))


	DB = database.InitDB()
	defer DB.Close()

	user,err:=services.GetUserByUsername("admin")
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("yaaaaaaaaaaaaaaaaaaaa",user)
	}

	// DB=database.GetDB()
	router := routes.SetupRouter()
	log.Println("Server running on port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))

}
