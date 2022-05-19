package controllerV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type auth struct {
	repository *repository.Repository
}

func newAuth(repository *repository.Repository) *auth {
	return &auth{repository: repository}
}

type authSignUpBody struct {
	Email    string `json:"email" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *auth) SignUp(ctx *gin.Context) {
	var body authSignUpBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, newWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	if c.repository.Users.Exists(&repository.UserModel{Email: body.Email}) {
		ctx.JSON(http.StatusBadRequest, newWrap(ErrEmailRegistred))
		return
	}

	model := repository.UserModel{
		Email:    body.Email,
		FullName: body.FullName,
		Password: body.Password,
	}

	if err := c.repository.Users.Create(&model); err != nil {
		ctx.JSON(http.StatusInternalServerError, newWrap(ErrServer.Add(err.Error())))
		return
	}

	ctx.JSON(http.StatusOK, newWrap(model.ID))
}
