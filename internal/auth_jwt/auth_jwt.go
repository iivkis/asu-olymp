package authjwt

import (
	"github.com/golang-jwt/jwt"
)

type AuthJWT struct {
	secret []byte
}

func NewAuthJWT(secret string) *AuthJWT {
	return &AuthJWT{
		secret: []byte(secret),
	}
}

func (j *AuthJWT) GenerateUserToken(claims *UserClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}
