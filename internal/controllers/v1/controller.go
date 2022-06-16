package ctrlv1

import (
	"github.com/iivkis/asu-olymp/internal/auth"
	"github.com/iivkis/asu-olymp/internal/repository"
)

//@Title ASU-Olymp API
//@Version 1.0-alpha
//@BasePath /api/v1
//@Host localhost:8080

//@Contanct.Name ivkis
//@Contact.Url https://t.me/iivkis

//@Accept json
//@Produce json

//@SecurityDefinitions.ApiKey ApiKey
//@In header
//@Name Authorization
//@Description JWT token for authorization

type Controller struct {
	Auth            *AuthController
	Tasks           *TasksController
	Questions       *QuestionsController
	Answers         *AnswersController
	CheckingAnswers *CheckingAnswersController

	Middleware *MiddlewareController
}

func NewController(repository *repository.Repository, authjwt *auth.AuthorizationJWT) *Controller {
	return &Controller{
		Auth:            NewAuthController(repository, authjwt),
		Tasks:           NewTasksController(repository),
		Questions:       NewQuestionsController(repository),
		Answers:         NewAnswersController(repository),
		CheckingAnswers: NewCheckingAnswersController(repository),
		Middleware:      NewMiddlewareController(authjwt),
	}
}
