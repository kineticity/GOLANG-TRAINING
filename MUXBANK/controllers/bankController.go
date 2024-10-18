package controllers

import (
	"bankingApp/models"
	"bankingApp/services"
	"bankingApp/validations"
	"encoding/json"
	"net/http"

)

func CreateBankController(w http.ResponseWriter, r *http.Request) {
	var bankRequestData models.Bank
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := validation.DecodeRequestBody(r, &bankRequestData); err != nil {
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
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	bankID, err := validation.GetIDFromRequest(r, "BankId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	banks := services.GetAllBanks()
	json.NewEncoder(w).Encode(banks)
}

func UpdateBankController(w http.ResponseWriter, r *http.Request) {
	var bankRequestData models.Bank
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	bankID, err := validation.GetIDFromRequest(r, "BankId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validation.DecodeRequestBody(r, &bankRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	bank, err := services.UpdateBank(bankID, bankRequestData.FullName, bankRequestData.Abbreviation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bank)
}

func DeleteBankController(w http.ResponseWriter, r *http.Request) {
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	bankID, err := validation.GetIDFromRequest(r, "BankId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.DeleteBank(bankID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
