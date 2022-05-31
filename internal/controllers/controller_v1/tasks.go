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

//@Summary Get tasks
//@Tags tasks
//@ID GetTasks
//@Success 200 {object} wrap{data=[]repository.TasksFindResult}
//@Failure 500
//@Router /t/tasks [get]
func (c *TasksController) Get(ctx *gin.Context) {
	models, err := c.repository.Tasks.Find(&repository.TaskModel{IsPublic: true}, getPayload(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}
	ctx.JSON(http.StatusOK, inWrap(models))
}

//@Summary Get one task by ID
//@Tags tasks
//@ID GetOneTask
//@Param id path int true "task ID"
//@Success 200 {object} wrap{data=repository.TasksFindResult}
//@Failure 400
//@Failure 404
//@Failure 500
//@Router /t/tasks/{id} [get]
func (c *TasksController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	model, err := c.repository.Tasks.FindByID(uint(id))
	if err != nil {
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
	Title       string `json:"title" binding:"required,max=200"`
	Content     string `json:"content" binding:"required,min=10,max=2000"`
	ShowCorrect bool   `json:"show_correct"`
	IsPublic    bool   `json:"is_public"`
}

//@Summary Create a new task
//@Security ApiKey
//@Tags tasks
//@ID AddTask
//@Param body body TasksPostBody true "task body"
//@Success 201 {object} wrap{data=DefaultOut}
//@Failure 400
//@Failure 500
//@Router /t/tasks [post]
func (c *TasksController) Post(ctx *gin.Context) {
	var body TasksPostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	model := repository.TaskModel{
		Title:       body.Title,
		Content:     body.Content,
		ShowCorrect: body.ShowCorrect,
		IsPublic:    body.IsPublic,
		AuthorID:    claims.ID,
	}

	if err := c.repository.Tasks.Cursor().Create(&model).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusCreated, inWrap(DefaultOut{ID: model.ID}))
}

type TasksPutBody struct {
	Title       *string `json:"title"`
	Content     *string `json:"content"`
	ShowCorrect *bool   `json:"show_correct"`
	IsPublic    *bool   `json:"is_public"`
}

//@Summary Update task fields
//@Security ApiKey
//@Tags tasks
//@ID UpdateTask
//@Param body body TasksPutBody true "task body"
//@Param id path int true "task ID"
//@Success 200 {object} wrap{data=DefaultOut}
//@Failure 400
//@Failure 404
//@Failure 500
//@Router /t/tasks/{id} [put]
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
		"title":        body.Title,
		"content":      body.Content,
		"show_correct": body.ShowCorrect,
		"is_public":    body.IsPublic,
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
		"show_correct": func(val interface{}) bool { return true },
		"is_public":    func(val interface{}) bool { return true },
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
