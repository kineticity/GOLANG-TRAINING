package controllers

import (
	"bankingApp/middlewares"
	"bankingApp/models"
	"bankingApp/services"
	validation "bankingApp/validations"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAccountController(w http.ResponseWriter, r *http.Request) {
	claims, err := validation.VerifyCustomerAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var accountRequestData models.Account

	if err := validation.DecodeRequestBody(r, &accountRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}


	fmt.Println("claims.userid:=",claims.UserID)

	var userCreatingAccount *models.User


	userCreatingAccount,err=services.GetCustomerByID(claims.UserID);if err!=nil{
		http.Error(w,"User Account not found",http.StatusBadRequest)
		return
	}

	//check if user already has bank account in this bankid bank
	for _,existingBankAccount:=range userCreatingAccount.Accounts{ //should maybe go in service?
		fmt.Println("ExistingBankaccount:",existingBankAccount)
		if accountRequestData.BankID==existingBankAccount.BankID{
			http.Error(w,"Customer already has an account in this bank",http.StatusBadRequest)
			return
		}
	}

	account, err := services.CreateAccount(claims.UserID, accountRequestData.BankID, accountRequestData.Balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(account)
}

func GetAccountByIDController(w http.ResponseWriter, r *http.Request) {
	claims, err := validation.VerifyCustomerAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}


	accountID, err := validation.GetIDFromRequest(r,"AccountId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	account, err := services.GetAccountByID(accountID)
	if err != nil || account.CustomerID != claims.UserID {
		http.Error(w, "Account not found or unauthorized access", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func GetAllAccountsController(w http.ResponseWriter, r *http.Request) {
	claims, err := validation.VerifyCustomerAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	accounts := services.GetAllAccounts(claims.UserID)
	json.NewEncoder(w).Encode(accounts)
}

func UpdateAccountController(w http.ResponseWriter, r *http.Request) {
	claims, err := validation.VerifyCustomerAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accountID, err := validation.GetIDFromRequest(r,"AccountId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var accountRequestData models.Account


	if err := validation.DecodeRequestBody(r, &accountRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	account, err := services.UpdateAccount(accountID, accountRequestData.BankID, accountRequestData.Balance)
	if err != nil || account.CustomerID != claims.UserID {
		http.Error(w, "Account not found or unauthorized access not your account", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func DeleteAccountController(w http.ResponseWriter, r *http.Request) {
	claims, err := validation.VerifyCustomerAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accountID, err := validation.GetIDFromRequest(r,"AccountId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	account, err := services.GetAccountByID(accountID)
	if err!=nil || account.CustomerID!=claims.UserID{
		http.Error(w,"Unauthorized Access Not your account",http.StatusUnauthorized)

	}

	if err := services.DeleteAccount(accountID,claims.UserID); err != nil  {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


func WithdrawController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsCustomer {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	accountIDStr := vars["AccountId"]

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		http.Error(w, "Invalid Account ID", http.StatusBadRequest)
		return
	}
	accountToWithdrawFrom,err:=services.GetAccountByID(accountID)
	if err!=nil{
		http.Error(w,"Account not found",http.StatusNotFound)
		return
	}
	if accountToWithdrawFrom.CustomerID!=claims.UserID{
		http.Error(w,"Account is not yours",http.StatusNotFound)
		return 
	}
		
	var requestData struct {
		Amount float64 `json:"amount"`
	}

	if err := validation.DecodeRequestBody(r, &requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	account, err := services.Withdraw(accountID, requestData.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func DepositController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsCustomer {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	accountIDStr := vars["AccountId"]

	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		http.Error(w, "Invalid Account ID", http.StatusBadRequest)
		return
	}

	accountToDepositTo,err:=services.GetAccountByID(accountID)
	if err!=nil{
		http.Error(w,"Account not found",http.StatusNotFound)
		return
	}
	if accountToDepositTo.CustomerID!=claims.UserID{
		http.Error(w,"Account is not yours",http.StatusNotFound)
		return 
	}

	var requestData struct {
		Amount float64 `json:"amount"`
	}

	if err := validation.DecodeRequestBody(r, &requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	account, err := services.Deposit(accountID, requestData.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func TransferController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsCustomer {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	fromAccountIDStr := vars["FromAccountId"]

	fromAccountID, err := strconv.Atoi(fromAccountIDStr)
	if err != nil {
		http.Error(w, "Invalid From Account ID", http.StatusBadRequest)
		return
	}

	accountToTransferFrom,err:=services.GetAccountByID(fromAccountID)
	if err!=nil{
		http.Error(w,"Account not found",http.StatusNotFound)
		return
	}
	if accountToTransferFrom.CustomerID!=claims.UserID{
		http.Error(w,"Account is not yours",http.StatusUnauthorized)
		return 
	}

	var requestData struct {
		ToAccountID int     `json:"toaccountid"`
		Amount      float64 `json:"amount"`
	}

	if err := validation.DecodeRequestBody(r, &requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if err := services.Transfer(fromAccountID, requestData.ToAccountID, requestData.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

