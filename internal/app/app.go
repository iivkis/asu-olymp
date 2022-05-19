package app

import (
	controllerV1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
	handlerHttp "github.com/iivkis/asu-olymp/internal/handlers/handler_http"
	"github.com/iivkis/asu-olymp/internal/repository"
)

func Launch() {
	repository := repository.NewRepository()

	cv1 := controllerV1.NewControllerV1(repository)

	handHttp := handlerHttp.NewHandlerHTTP(&handlerHttp.Config{
		ControllerV1: cv1,
	})

	handHttp.Engine().Run(":8081")
}
