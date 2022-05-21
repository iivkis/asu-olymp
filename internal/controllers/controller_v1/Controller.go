package controllerV1

import (
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	"github.com/iivkis/asu-olymp/internal/repository"
)

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
