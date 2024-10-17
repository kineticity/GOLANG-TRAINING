package services

import (
	// "bankingApp/models"
	// "errors"
	// "time"

	// "github.com/dgrijalva/jwt-go"
)

// var users = make(map[string]*models.UserToken) 



// var Userslist = []*models.User{
// 	{UserID: 1, Username: "admin", Password: "adminpass", IsAdmin: true, IsCustomer: false},
// 	{UserID: 2, Username: "customer", Password: "customerpass", IsAdmin: false, IsCustomer: true},
// }

// ValidateUser checks if the user exists and if the password is correct
// func ValidateUser(username, password string) (*models.User, error) {
// 	for _, user := range Userslist {
// 		if user.Username == username {
// 			if user.Password == password {
// 				return user, nil 
// 			}
// 			return nil, errors.New("incorrect password")
// 		}
// 	}
// 	return nil, errors.New("user not found")
// }
// func GenerateToken(username string, isAdmin bool) (string, error) {
// 	claims := &models.Claims{
// 		Username:  username,
// 		IsAdmin:   isAdmin,
// 		ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(models.SecretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	// store token and user info in users map
// 	users[username] = &models.UserToken{
// 		Username:  username,
// 		Token:     tokenString,
// 		ExpiresAt: time.Now().Add(3* time.Hour),
// 	}

// 	return tokenString, nil
// }

// func VerifyUserToken(username, token string) (*models.Claims, error) {
// 	user, exists := users[username]
// 	if !exists || user.Token != token {
// 		return nil, errors.New("invalid token or user")
// 	}

// 	if time.Now().After(user.ExpiresAt) { //not checked
// 		return nil, errors.New("token has expired")
// 	}

// 	claims := &models.Claims{}
// 	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return claims, nil
// }
