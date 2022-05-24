package controllerV1

import (
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	"github.com/iivkis/asu-olymp/internal/repository"
)

//@Title ASU-Olymp API
//@Version 1.0-alpha
//@BasePath /api/v1
//@Host localhost:8081

//@Contanct.Name ivkis
//@Contact.Url https://t.me/iivkis

//@Accept json
//@Produce json

//@SecurityDefinitions.ApiKey ApiKey
//@In header
//@Name Authorization
//@Description JWT token for authorization

type ControllerV1 struct {
	Auth      *AuthController
	Tasks     *TasksController
	Questions *QuestionsController
	Answers   *AnswersController

	Middleware *MiddlewareController
}

func NewControllerV1(repository *repository.Repository, authjwt *authjwt.AuthJWT) *ControllerV1 {
	return &ControllerV1{
		Auth:       NewAuthController(repository, authjwt),
		Tasks:      NewTasksController(repository),
		Questions:  NewQuestionsController(repository),
		Answers:    NewAnswersController(repository),
		Middleware: NewMiddlewareController(authjwt),
	}
}
