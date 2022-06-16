package app

import (
	"fmt"

	"github.com/iivkis/asu-olymp/config"
	"github.com/iivkis/asu-olymp/internal/auth"
	ctrlv1 "github.com/iivkis/asu-olymp/internal/controllers/v1"
	httphr "github.com/iivkis/asu-olymp/internal/handlers/http"
	"github.com/iivkis/asu-olymp/internal/repository"
)

func Launch() {
	authjwt := auth.NewAuthorizationWithJWT(config.JWT_SECRET)

	repository := repository.NewRepository()

	cv1 := ctrlv1.NewController(repository, authjwt)

	httpHandler := httphr.NewHandlerHTTP(&httphr.Config{
		ControllerV1: cv1,
	})

	httpHandler.Engine().Run(fmt.Sprintf(":%s", config.PORT))
}
