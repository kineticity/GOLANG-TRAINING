package controllers

import (
	"bankingApp/models"
	"bankingApp/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Username==""||user.Password==""||user.FirstName==""||user.LastName==""{
		w.WriteHeader(http.StatusBadRequest)
	}

	if _,err := services.CreateCustomer(user.Username, user.Password, user.FirstName, user.LastName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id<=0{
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}


	if err := services.DeleteCustomerByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id<=0{
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	customer,err := services.GetCustomerByID(id); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(customer)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
    idStr := mux.Vars(r)["id"]
    id, err := strconv.Atoi(idStr)
    if err != nil ||id<=0{
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var updateData models.User

    if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	if updateData.Username==""||updateData.Password==""||updateData.FirstName==""||updateData.LastName==""{
		http.Error(w,"invalid request body",http.StatusBadRequest)
	}

    user,err := services.UpdateCustomerByID(id, updateData.FirstName, updateData.LastName, updateData.Username, updateData.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}


func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}
