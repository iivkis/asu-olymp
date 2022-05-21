package controllerV1

import (
	"net/http"
	"strconv"

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
	TaskID uint `form:"task_id" binding:"min=0"`
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

type QuestionsPutBody struct {
	Text *string `json:"text"`
}

func (c *QuestionsController) Put(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	claims, _ := getUserClaims(ctx)

	if !c.repository.Questions.Exists(&repository.QuestionModel{ID: uint(id), AuthorID: claims.ID}) {
		ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound))
		return
	}

	var body QuestionsPutBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	fields := map[string]interface{}{
		"text": body.Text,
	}

	if err := validator(fields, validatorRules{
		"text": func(val interface{}) bool {
			l := len(*val.(*string))
			return l > 1 && l < 1000
		},
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	if err := c.repository.Questions.Update(&repository.QuestionModel{ID: uint(id)}, fields); err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	ctx.JSON(http.StatusOK, inWrap(DefaultOut{ID: uint(id)}))
}
