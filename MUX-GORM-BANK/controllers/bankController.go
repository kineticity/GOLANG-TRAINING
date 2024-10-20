package controllers

import (
	"bankingApp/models"
	"bankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateBank(w http.ResponseWriter, r *http.Request) {
	var bank models.Bank
	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := services.CreateBank(bank.FullName, bank.Abbreviation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteBankByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.DeleteBankByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBankByID(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid bank ID", http.StatusBadRequest)
        return
    }

    var updatedBank models.Bank
    if err := json.NewDecoder(r.Body).Decode(&updatedBank); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    bank,err := services.UpdateBankByID(id, updatedBank.FullName, updatedBank.Abbreviation)
    if err != nil {
        
        http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
        
    

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bank)
    
}


func GetAllBanks(w http.ResponseWriter, r *http.Request) {
	banks, err := services.GetAllBanks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(banks)
}

func GetBankByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	customer,err := services.GetBankByID(uint(id)); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customer)
}
