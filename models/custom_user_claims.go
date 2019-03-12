package models

import "github.com/dgrijalva/jwt-go"

type JwtCustomUserClaims struct {
	ID int64 `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}