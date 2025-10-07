package domain

import "github.com/golang-jwt/jwt/v5"

type JWT struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
