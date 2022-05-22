package controllerV1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
	"gorm.io/gorm"
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

func (c *AnswersController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	var model repository.AnswerModel
	if err := c.repository.Answers.Cursor().First(&model, &repository.AnswerModel{ID: uint(id), AuthorID: claims.ID}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
		} else {
			ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		}
		return
	}

	ctx.JSON(http.StatusOK, inWrap(model))
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

type AnswersPutBody struct {
	Value *string `json:"value"`
}

func (c *AnswersController) Put(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	if !c.repository.Answers.Exists(&repository.AnswerModel{ID: uint(id), AuthorID: claims.ID}) {
		ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
		return
	}

	var body AnswersPutBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	fields := map[string]interface{}{
		"value": body.Value,
	}

	if err := validator(fields, validatorRules{
		"value": func(val interface{}) bool {
			l := len(*val.(*string))
			return l >= 1 && l <= 1000
		},
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	if err := c.repository.Answers.Update(&repository.AnswerModel{ID: uint(id)}, fields); err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusOK, inWrap(DefaultOut{ID: uint(id)}))
}
