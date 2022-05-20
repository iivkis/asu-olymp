package authjwt

import (
	"fmt"

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

func (j *AuthJWT) ParseUserToken(t string) (claims *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(t, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(2)
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*UserClaims), nil
}
