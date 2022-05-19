package controllerV1

import (
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type ControllerV1 struct {
	Auth *AuthController
}

func NewControllerV1(repository *repository.Repository, authjwt *authjwt.AuthJWT) *ControllerV1 {
	return &ControllerV1{
		Auth: NewAuthController(repository, authjwt),
	}
}
