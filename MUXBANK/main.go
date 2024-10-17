package main

import (
	"bankingApp/controllers"
	"bankingApp/middlewares"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Authentication
	router.HandleFunc("/login", controllers.LoginController).Methods(http.MethodPost)

	// Authorize admin Routes
	adminRoutes := router.PathPrefix("/adminuser").Subrouter()
	adminRoutes.Use(middlewares.TokenAuthMiddleware) //verify jwt
	adminRoutes.Use(middlewares.VerifyAdmin) // Authorize if admin

	adminRoutes.HandleFunc("/customer", controllers.CreateCustomerController).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/customer/{UserId}", controllers.GetCustomerByIDController).Methods(http.MethodGet)
	adminRoutes.HandleFunc("/customer", controllers.GetAllCustomersController).Methods(http.MethodGet)
	adminRoutes.HandleFunc("/customer/{UserId}", controllers.UpdateCustomerController).Methods(http.MethodPut)
	adminRoutes.HandleFunc("/customer/{UserId}", controllers.DeleteCustomerController).Methods(http.MethodDelete)

	adminRoutes.HandleFunc("/bank", controllers.CreateBankController).Methods(http.MethodPost)
	adminRoutes.HandleFunc("/bank/{BankId}", controllers.GetBankByIDController).Methods(http.MethodGet)
	adminRoutes.HandleFunc("/bank", controllers.GetAllBanksController).Methods(http.MethodGet)
	adminRoutes.HandleFunc("/bank/{BankId}", controllers.UpdateBankController).Methods(http.MethodPut)
	adminRoutes.HandleFunc("/bank/{BankId}", controllers.DeleteBankController).Methods(http.MethodDelete)

	// adminRoutes.HandleFunc("bank/ledger", controllers.AddToLedgerController).Methods(http.MethodPost)
	// adminRoutes.HandleFunc("bank/{BankId}/ledger", controllers.GetLedgerController).Methods(http.MethodGet)

	adminRoutes.HandleFunc("/ledger/add", controllers.AddToLedgerController).Methods("POST")
    adminRoutes.HandleFunc("/ledger/{BankId}", controllers.GetLedgerController).Methods("GET")





	// Authorize customer routes
	customerRoutes := router.PathPrefix("/customeruser").Subrouter()
	customerRoutes.Use(middlewares.TokenAuthMiddleware) // Verify JWT
	customerRoutes.Use(middlewares.VerifyCustomer) // Authorize if customer

	customerRoutes.HandleFunc("/accounts", controllers.CreateAccountController).Methods(http.MethodPost)
	customerRoutes.HandleFunc("/accounts/{AccountId}", controllers.GetAccountByIDController).Methods(http.MethodGet)
	customerRoutes.HandleFunc("/accounts", controllers.GetAllAccountsController).Methods(http.MethodGet)
	customerRoutes.HandleFunc("/accounts/{AccountId}", controllers.UpdateAccountController).Methods(http.MethodPut)
	customerRoutes.HandleFunc("/accounts/{AccountId}", controllers.DeleteAccountController).Methods(http.MethodDelete)

	customerRoutes.HandleFunc("/accounts/{AccountId}/withdraw", controllers.WithdrawController).Methods(http.MethodPost)
	customerRoutes.HandleFunc("/accounts/{AccountId}/deposit", controllers.DepositController).Methods(http.MethodPost)
	customerRoutes.HandleFunc("/accounts/{FromAccountId}/transfer", controllers.TransferController).Methods(http.MethodPost)

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
