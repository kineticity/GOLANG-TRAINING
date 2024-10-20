package models

import (
	"errors"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/golang-jwt/jwt"
)

var SecretKey = []byte("secret_key")

type UserToken struct {
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
	UserID    uint      `json:"userId"`
}

type Claims struct {
	Username   string `json:"username"`
	IsAdmin    bool   `json:"isAdmin"`
	IsCustomer bool   `json:"isCustomer"`
	ExpiresAt  int64  `json:"exp"` // expiration time
	UserID     uint   `json:"userId"`
}

func NewClaims(username string, isAdmin, isCustomer bool, expirationDate time.Time, userid uint) *Claims {
	return &Claims{
		Username:   username,
		IsAdmin:    isAdmin,
		IsCustomer: isCustomer,
		ExpiresAt:  expirationDate.Unix(),
		UserID:     userid,
	}
}

func (c *Claims) Valid() error {
	if time.Now().Unix() > c.ExpiresAt {
		return errors.New("token has expired")
	}
	return nil
}

func (c *Claims) Signing() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(SecretKey)
}
