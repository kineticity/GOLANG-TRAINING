package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"bankingApp/models"
	"bankingApp/services"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var users = make(map[int]*models.UserToken)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) 
	return err == nil
}

func ValidateUser(username, password string) (*models.User, error) {
	user, err := services.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !CheckPasswordHash(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials) 
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := ValidateUser(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	var claim *models.Claims
	if user.IsAdmin && !user.IsCustomer {
		claim = models.NewClaims(user.Username, true, false, time.Now().Add(time.Hour*3), user.ID)
	} else if !user.IsAdmin && user.IsCustomer {
		claim = models.NewClaims(user.Username, false, true, time.Now().Add(time.Hour*3), user.ID)
	} else {
		http.Error(w, "User has some role conflicts", http.StatusUnauthorized)
		return
	}

	tokenString, err := claim.Signing()
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	tokenUser := &models.UserToken{UserID: user.ID, Token: tokenString, Username: user.Username}
	users[int(user.ID)] = tokenUser

	w.Header().Set("Authorization", tokenString)
	json.NewEncoder(w).Encode(claim)
}
