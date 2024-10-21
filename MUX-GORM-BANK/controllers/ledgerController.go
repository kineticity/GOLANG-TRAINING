package controllers

import (
	"bankingApp/models"
	"bankingApp/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateLedgerEntry(w http.ResponseWriter, r *http.Request) {
	var entry models.LedgerEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if entry.Amount<=0||entry.EntryType==""||entry.BankID<=0||entry.CorrespondingBankID<=0{
		http.Error(w,"invalid request body",http.StatusBadRequest)
	}

	fmt.Println("entrytype:",entry.EntryType)

	if _, err := services.CreateLedgerEntry(entry.BankID,entry.CorrespondingBankID, entry.Amount,entry.EntryType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetLedgerEntries(w http.ResponseWriter, r *http.Request) {
	bankIDStr := mux.Vars(r)["bankID"]
	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil ||bankID<=0{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ledgerEntries, err := services.GetLedgerEntries(bankID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ledgerEntries)
}
