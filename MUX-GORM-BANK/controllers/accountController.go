package controllers

import (
	"bankingApp/middlewares"
	"bankingApp/models"
	"bankingApp/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	customerID := claims.UserID

	var account models.Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account.CustomerID = customerID

	if _, err := services.CreateAccount(account.CustomerID, account.BankID, account.Balance); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteAccountByID(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := mux.Vars(r)["id"]
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteAccountByID(accountID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	accounts, err := services.GetAllAccountsByUserID(claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(accounts)
}
func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := mux.Vars(r)["id"]
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := services.GetAccountByID(claims.UserID, uint(accountID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}
func UpdateAccountByID(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := mux.Vars(r)["id"]
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var updatedAccount models.Account
	if err := json.NewDecoder(r.Body).Decode(&updatedAccount); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedAccount.CustomerID = claims.UserID

	if err := services.UpdateAccountByID(claims.UserID, uint(accountID), updatedAccount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}


func WithdrawFromAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := services.Withdraw(id, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func DepositToAccount(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := services.Deposit(id, request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(account)
}

func TransferBetweenAccounts(w http.ResponseWriter, r *http.Request) {
	fromAccountIDStr := mux.Vars(r)["fromAccountID"]
	fromAccountID, err := strconv.Atoi(fromAccountIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	var request struct {
		Amount float64 `json:"amount"`
		ToAccountID int `json:"toAccountID"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.Transfer(fromAccountID, request.ToAccountID, request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PrintPassbookController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsCustomer {
		http.Error(w, "Unauthorized: Only customers can view their passbook", http.StatusUnauthorized)
		return
	}

	fmt.Println("before idstr")
	idStr := mux.Vars(r)["id"]
	fmt.Println("idstr:",idStr)
	accountID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("aftre error idstr")

		return
	}
	fmt.Println("after no error idstr")


	err = services.PrintAccountPassbook(accountID)
	if err != nil {
		http.Error(w, "Failed to print passbook: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Passbook printed successfully"))
}
