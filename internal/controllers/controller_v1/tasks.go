package controllerV1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
	"gorm.io/gorm"
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
	models, err := c.repository.Tasks.Find(&repository.TaskModel{}, getPayload(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}
	ctx.JSON(http.StatusOK, inWrap(models))
}

func (c *TasksController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	var model *repository.TaskModel
	if err := c.repository.Tasks.Cursor().First(&model, &repository.TaskModel{ID: uint(id)}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
		} else {
			ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		}
		return
	}
	ctx.JSON(http.StatusOK, inWrap(model))
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

	if err := c.repository.Tasks.Cursor().Create(&model); err != nil {
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

	claims, _ := getUserClaims(ctx)

	if !c.repository.Tasks.Exists(&repository.TaskModel{ID: uint(id), AuthorID: claims.ID}) {
		ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
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
			l := len(*val.(*string))
			return l > 1 && l < 200
		},
		"content": func(val interface{}) bool {
			l := len(*val.(*string))
			return l > 1 && l < 2000
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
