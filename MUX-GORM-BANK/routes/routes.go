package routes

import (
	"bankingApp/controllers"
	"bankingApp/middlewares"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.LoginController).Methods(http.MethodPost)

	adminRoutes := router.PathPrefix("/adminuser").Subrouter()
	adminRoutes.Use(middlewares.TokenAuthMiddleware) 
	adminRoutes.Use(middlewares.VerifyAdmin)

	adminRoutes.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	adminRoutes.HandleFunc("/users/{id}", controllers.DeleteUserByID).Methods("DELETE")
	adminRoutes.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	adminRoutes.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	adminRoutes.HandleFunc("/users/{id}", controllers.UpdateUserByID).Methods("PUT")

	adminRoutes.HandleFunc("/banks", controllers.CreateBank).Methods("POST")
	adminRoutes.HandleFunc("/banks/{id}", controllers.DeleteBankByID).Methods("DELETE")
	adminRoutes.HandleFunc("/banks", controllers.GetAllBanks).Methods("GET")
	adminRoutes.HandleFunc("/banks/{id}", controllers.GetBankByID).Methods("GET")
	adminRoutes.HandleFunc("/banks/{id}", controllers.UpdateBankByID).Methods("PUT")

	adminRoutes.HandleFunc("/ledger", controllers.CreateLedgerEntry).Methods("POST")
	adminRoutes.HandleFunc("/ledger/{bankID}", controllers.GetLedgerEntries).Methods("GET")

	
	customerRoutes := router.PathPrefix("/customeruser").Subrouter()
	customerRoutes.Use(middlewares.TokenAuthMiddleware) 
	customerRoutes.Use(middlewares.VerifyCustomer)      

	customerRoutes.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")
	customerRoutes.HandleFunc("/accounts", controllers.GetAllAccounts).Methods("GET")
	customerRoutes.HandleFunc("/accounts/{id}", controllers.GetAccountByID).Methods("GET")
	customerRoutes.HandleFunc("/accounts/{id}", controllers.DeleteAccountByID).Methods("DELETE")
	customerRoutes.HandleFunc("/accounts/{id}", controllers.UpdateAccountByID).Methods("PUT")

	customerRoutes.HandleFunc("/accounts/{id}/withdraw", controllers.WithdrawFromAccount).Methods("POST")
	customerRoutes.HandleFunc("/accounts/{id}/deposit", controllers.DepositToAccount).Methods("POST")
	customerRoutes.HandleFunc("/accounts/{fromAccountID}/transfer", controllers.TransferBetweenAccounts).Methods("POST")
	customerRoutes.HandleFunc("/accounts/{id}/passbook", controllers.PrintPassbookController).Methods("GET")


	return router
}
