package controllerV1

import "github.com/iivkis/asu-olymp/internal/repository"

type ControllerV1 struct {
	Ping *ping
	Auth *auth
}

func NewControllerV1(repository *repository.Repository) *ControllerV1 {
	return &ControllerV1{
		Ping: newPing(),
		Auth: newAuth(repository),
	}
}
