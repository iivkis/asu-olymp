package ctrlv1

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
	"gorm.io/gorm"
)

type CheckingAnswersController struct {
	repository *repository.Repository
}

func NewCheckingAnswersController(repository *repository.Repository) *CheckingAnswersController {
	return &CheckingAnswersController{
		repository: repository,
	}
}

type CheckingAnswersPostBodyFieldAnswer struct {
	QuestionID uint   `json:"question_id" binding:"min=1"`
	Value      string `json:"value" binding:"max=1000"`
}

type CheckingAnswersPostBody struct {
	TaskID  uint                                 `json:"task_id" binding:"min=1"`
	Answers []CheckingAnswersPostBodyFieldAnswer `json:"answers"`
}

type CheckingAnswersPostOut struct {
	TaskID       uint          `json:"task_id"`
	NumOfCorrect uint          `json:"num_of_correct"`
	ShowCorrect  bool          `json:"show_correct"`
	Results      map[uint]bool `json:"results"`
}

//@Summary Check of correct answers
//@Security ApiKey
//@Tags checking
//@ID CheckingAnswers
//@Param body body CheckingAnswersPostBody true "check answers body"
//@Success 200 {object} wrap{data=CheckingAnswersPostOut}
//@Failure 400
//@Failure 500
//@Router /t/check [post]
func (c *CheckingAnswersController) Post(ctx *gin.Context) {
	var body CheckingAnswersPostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	var task repository.TaskModel
	if err := c.repository.Tasks.Cursor().Select("show_correct").Where("id = ?", body.TaskID).First(&task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, inWrap(ErrRecordNotFound.Add("not found task with this ID")))
		} else {
			ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		}
		return
	}

	questExistence, err := c.repository.Questions.MapOfExistence(&repository.QuestionModel{TaskID: body.TaskID})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	output := CheckingAnswersPostOut{TaskID: body.TaskID}
	if task.ShowCorrect {
		output.ShowCorrect = true
		output.Results = make(map[uint]bool)
	}

	needQuestID := make([]uint, 0, len(questExistence))
	for _, usrAnswer := range body.Answers {
		if _, ok := questExistence[usrAnswer.QuestionID]; ok {
			needQuestID = append(needQuestID, usrAnswer.QuestionID)
		} else {
			msg := fmt.Sprintf("undefined `question_id: %d` in `task_id: %d`", usrAnswer.QuestionID, body.TaskID)
			ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(msg)))
			return
		}
	}

	originAnswersMap, err := c.repository.Answers.FindAndTransformToMapByQuestionID(needQuestID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer))
		return
	}

	for _, usrAnswer := range body.Answers {
		originAnswer, ok := originAnswersMap[usrAnswer.QuestionID]

		eq := !ok || strings.EqualFold(strings.TrimSpace(usrAnswer.Value), originAnswer.Value)
		if eq {
			output.NumOfCorrect += 1
		}

		if output.ShowCorrect {
			output.Results[usrAnswer.QuestionID] = eq
		}
	}

	ctx.JSON(http.StatusOK, inWrap(output))
}
