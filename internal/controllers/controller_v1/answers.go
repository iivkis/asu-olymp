package controllerV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type AnswersController struct {
	repository *repository.Repository
}

func NewAnswersController(repository *repository.Repository) *AnswersController {
	return &AnswersController{repository: repository}
}

type AnswersGetQuery struct {
	QuestionID uint `form:"question_id" binding:"min=0"`
}

func (c *AnswersController) Get(ctx *gin.Context) {
	var query AnswersGetQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	models, err := c.repository.Answers.Find(&repository.AnswerModel{AuthorID: claims.ID, QuestionID: query.QuestionID}, getPayload(ctx))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusOK, inWrap(models))
}

type AnswerPostBody struct {
	Value      string `json:"value" binding:"required,max=1000"`
	QuestionID uint   `json:"question_id" binding:"required,min=1"`
}

func (c *AnswersController) Post(ctx *gin.Context) {
	var body AnswerPostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	if !c.repository.Questions.Exists(&repository.QuestionModel{ID: body.QuestionID, AuthorID: claims.ID}) {
		ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
		return
	}

	model := repository.AnswerModel{
		Value:      body.Value,
		QuestionID: body.QuestionID,
		AuthorID:   claims.ID,
	}

	if err := c.repository.Answers.Cursor().Create(&model).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusCreated, inWrap(DefaultOut{ID: model.ID}))
}
