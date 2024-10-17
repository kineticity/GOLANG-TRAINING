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


var customerRequestData models.User

func CreateCustomerController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization")) //dk if need to check again
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&customerRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	fmt.Println(customerRequestData)

	if customerRequestData.FirstName==""||customerRequestData.LastName==""||customerRequestData.Password==""||customerRequestData.Username==""{
		http.Error(w,"Compulsary fields are empty",http.StatusBadRequest)
		return
	}

	customer, err := services.CreateCustomer(customerRequestData.Username, customerRequestData.Password, customerRequestData.FirstName, customerRequestData.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	json.NewEncoder(w).Encode(customer)
}

func GetCustomerByIDController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["UserId"]

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid UserId parameter", http.StatusBadRequest)
		return
	}

	customer, err := services.GetCustomerByID(userID)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func GetAllCustomersController(w http.ResponseWriter, r *http.Request) {
	customers := services.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

func UpdateCustomerController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["UserId"]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}
	existingCustomer, err := services.GetCustomerByID(userID)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&customerRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	if customerRequestData.Username == "" {
		customerRequestData.Username = existingCustomer.Username
	}
	if customerRequestData.Password == "" {
		customerRequestData.Password = existingCustomer.Password
	}
	if customerRequestData.FirstName == "" {
		customerRequestData.FirstName = existingCustomer.FirstName
	}
	if customerRequestData.LastName == "" {
		customerRequestData.LastName = existingCustomer.LastName
	}


	customer, err := services.UpdateCustomer(userID, customerRequestData.Username, customerRequestData.Password, customerRequestData.FirstName, customerRequestData.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func DeleteCustomerController(w http.ResponseWriter, r *http.Request) {
	claims, err := middlewares.VerifyJWT(r.Header.Get("Authorization"))
	if err != nil || !claims.IsAdmin {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["UserId"]
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteCustomer(userID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
