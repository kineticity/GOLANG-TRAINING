package controllers

import (
	"bankingApp/middlewares"
	"bankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var bankRequestData struct { //change to bank struct and check
	FullName     string `json:"fullname"`
	Abbreviation string `json:"abbreviation"`
}

func CreateBankController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&bankRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	bank, err := services.CreateBank(bankRequestData.FullName, bankRequestData.Abbreviation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func GetBankByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["BankId"]

	bankID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid BankId parameter", http.StatusBadRequest)
		return
	}

	bank, err := services.GetBankByID(bankID)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func GetAllBanksController(w http.ResponseWriter, r *http.Request) {
	banks := services.GetAllBanks()
	json.NewEncoder(w).Encode(banks)
}

func UpdateBankController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["BankId"]
	bankID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Bank ID", http.StatusBadRequest)
		return
	}
	existingBank, err := services.GetBankByID(bankID)
	if err != nil {
		http.Error(w, "Bank not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&bankRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if bankRequestData.FullName == "" {
		bankRequestData.FullName = existingBank.FullName
	}
	if bankRequestData.Abbreviation == "" {
		bankRequestData.Abbreviation = existingBank.Abbreviation
	}

	bank, err := services.UpdateBank(bankID, bankRequestData.FullName, bankRequestData.Abbreviation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func DeleteBankController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["BankId"]
	bankID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Bank ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteBank(bankID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

