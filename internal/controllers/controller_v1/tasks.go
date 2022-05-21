package controllerV1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type TaskOut struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type TasksController struct {
	repository *repository.Repository
}

func NewTasksController(repository *repository.Repository) *TasksController {
	return &TasksController{
		repository: repository,
	}
}

func (c *TasksController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, inWrap("ok"))
}

type TasksPostBody struct {
	Title   string `json:"title" binding:"required,max=200"`
	Content string `json:"content" binding:"required,min=10,max=2000"`
}

func (c *TasksController) Post(ctx *gin.Context) {
	var body TasksPostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	model := repository.TaskModel{
		Title:    body.Title,
		Content:  body.Content,
		AuthorID: claims.ID,
	}

	if err := c.repository.Tasks.CreateTask(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusCreated, inWrap(DefaultOut{ID: model.ID}))
}

type TasksPutBody struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (c *TasksController) Put(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	var body TasksPutBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	fields := map[string]interface{}{
		"title":   body.Title,
		"content": body.Content,
	}

	if err := validator(fields, validatorRules{
		"title": func(val interface{}) bool {
			str := *val.(*string)
			return len(str) > 1 && len(str) < 200
		},
		"content": func(val interface{}) bool {
			str := *val.(*string)
			return len(str) > 1 && len(str) < 2000
		},
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	if err := c.repository.Tasks.Update(&repository.TaskModel{ID: uint(id)}, fields); err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusOK, inWrap(DefaultOut{ID: uint(id)}))
}
