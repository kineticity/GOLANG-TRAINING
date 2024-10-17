package controllers

import (
	"bankingApp/middlewares"
	"bankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddToLedgerController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var requestData struct {
		LendingBankID    int     `json:"lendingbankid"`
		ReceivingBankID  int     `json:"receivingbankid"`
		Amount           float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if err := services.AddToLedger(requestData.LendingBankID, requestData.ReceivingBankID, requestData.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetLedgerController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bankIDStr := vars["BankId"]

	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		http.Error(w, "Invalid Bank ID", http.StatusBadRequest)
		return
	}

	ledger, err := services.GetLedger(bankID)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(ledger)
}

