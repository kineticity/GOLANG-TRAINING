package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"bankingApp/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)


var users = make(map[int]*models.UserToken)



func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) //if password and hash same returns nil
	return err == nil 
}

func ValidateUser(username, password string) (*models.User, error) {
	for _, user := range models.Userslist {
		if user.Username == username {
			if CheckPasswordHash(password, user.Password) {
				return user, nil
			}
			return nil, errors.New("incorrect password")
		}
	}
	return nil, errors.New("user not found")
}


func LoginController(w http.ResponseWriter, r *http.Request) {

	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials) //unmarshall
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := ValidateUser(credentials.Username, credentials.Password) //validate
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	var claim *models.Claims
	if user.IsAdmin && !user.IsCustomer {
		claim = models.NewClaims(user.Username, true, false, time.Now().Add(time.Hour*3), user.UserID)
	} else if !user.IsAdmin && user.IsCustomer {
		claim = models.NewClaims(user.Username, false, true, time.Now().Add(time.Hour*3), user.UserID)
	} else {
		http.Error(w, "User has some role conflicts", http.StatusUnauthorized)
		return
	}

	tokenString, err := claim.Signing() //signing()->generate jwt
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	tokenUser := &models.UserToken{UserID: user.UserID, Token: tokenString, Username: user.Username} //store token info
	users[user.UserID] = tokenUser

	w.Header().Set("Authorization", tokenString)

	json.NewEncoder(w).Encode(claim)
}
