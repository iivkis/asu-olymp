package auth

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}
