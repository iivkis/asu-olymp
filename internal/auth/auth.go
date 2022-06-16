package auth

import (
	"github.com/golang-jwt/jwt"
)

type AuthorizationJWT struct {
	secret []byte
}

func NewAuthorizationWithJWT(secret string) *AuthorizationJWT {
	return &AuthorizationJWT{
		secret: []byte(secret),
	}
}

func (j *AuthorizationJWT) GenerateUserToken(claims *UserClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
}

func (j *AuthorizationJWT) ParseUserToken(t string) (claims *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(t, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*UserClaims), nil
}
