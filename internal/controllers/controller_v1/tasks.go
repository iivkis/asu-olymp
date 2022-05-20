package controllerV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type TasksController struct {
	repository *repository.Repository
}

func NewTasksController(repository *repository.Repository) *TasksController {
	return &TasksController{
		repository: repository,
	}
}

func (c *TasksController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, newWrap("ok"))
}
