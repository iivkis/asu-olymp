package controllerV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type QuestionsController struct {
	repository *repository.Repository
}

func NewQuestionsController(repository *repository.Repository) *QuestionsController {
	return &QuestionsController{repository: repository}
}

type QuestionsGetQuery struct {
	TaskID uint `json:"task_id" binding:"min=0"`
}

func (c *QuestionsController) Get(ctx *gin.Context) {
	var query QuestionsGetQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	models, err := c.repository.Questions.Find(&repository.QuestionModel{TaskID: query.TaskID}, getPayload(ctx))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	ctx.JSON(http.StatusOK, inWrap(models))
}

type QuestionsPostBody struct {
	Text   string `json:"text" binding:"required,max=1000"`
	TaskID uint   `json:"task_id" binding:"min=1"`
}

func (c *QuestionsController) Post(ctx *gin.Context) {
	var body QuestionsPostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	if !c.repository.Tasks.Exists(&repository.TaskModel{ID: body.TaskID, AuthorID: claims.ID}) {
		ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
		return
	}

	model := repository.QuestionModel{
		Text:     body.Text,
		TaskID:   body.TaskID,
		AuthorID: claims.ID,
	}

	if err := c.repository.Questions.Cursor().Create(&model).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusCreated, inWrap(DefaultOut{ID: model.ID}))
}
