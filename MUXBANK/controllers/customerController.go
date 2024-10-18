package controllers

import (
	"bankingApp/models"
	"bankingApp/services"
	"bankingApp/validations"
	"encoding/json"
	"net/http"
)



func CreateCustomerController(w http.ResponseWriter, r *http.Request) {

	var customerRequestData models.User

	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := validation.DecodeRequestBody(r, &customerRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

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

	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := validation.GetIDFromRequest(r,"UserId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	customers := services.GetAllCustomers()
	json.NewEncoder(w).Encode(customers)
}

func UpdateCustomerController(w http.ResponseWriter, r *http.Request) {

	var customerRequestData models.User

	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	userID, err := validation.GetIDFromRequest(r,"UserId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validation.DecodeRequestBody(r, &customerRequestData); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	customer, err := services.UpdateCustomer(userID, customerRequestData.Username, customerRequestData.Password, customerRequestData.FirstName, customerRequestData.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func DeleteCustomerController(w http.ResponseWriter, r *http.Request) {
	
	_, err := validation.VerifyAdminAuthorization(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, err := validation.GetIDFromRequest(r,"UserId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.DeleteCustomer(userID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
