package app

import (
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	controllerV1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
	handlerHttp "github.com/iivkis/asu-olymp/internal/handlers/handler_http"
	"github.com/iivkis/asu-olymp/internal/repository"
)

func Launch() {

	authjwt := authjwt.NewAuthJWT("AllYourBase")
	repository := repository.NewRepository()

	cv1 := controllerV1.NewControllerV1(repository, authjwt)

	handHttp := handlerHttp.NewHandlerHTTP(&handlerHttp.Config{
		ControllerV1: cv1,
	})

	handHttp.Engine().Run(":8081")
}
