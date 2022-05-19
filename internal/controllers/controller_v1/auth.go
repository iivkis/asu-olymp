package controllerV1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type AuthController struct {
	repository *repository.Repository
	authjwt    *authjwt.AuthJWT
}

func NewAuthController(repository *repository.Repository, authjwt *authjwt.AuthJWT) *AuthController {
	return &AuthController{repository: repository, authjwt: authjwt}
}

type authSignUpBody struct {
	Email    string `json:"email" binding:"required,min=3,max=100"`
	FullName string `json:"full_name" binding:"required,min=1,max=100"`
	Password string `json:"password" binding:"required,min=4,max=50"`
}

func (c *AuthController) SignUp(ctx *gin.Context) {
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

type authSignInBody struct {
	Email    string `json:"email" binding:"required,min=3,max=200"`
	Password string `json:"password" binding:"required,min=4,max=50"`
}

type authSignInOut struct {
	Token string `json:"token"`
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var body authSignInBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, newWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	user, err := c.repository.Users.SignUpByEmail(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newWrap(ErrIncorrectData))
		return
	}

	claims := authjwt.UserClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 1000,
			Issuer:    "asu-olymp",
		},
	}
	token, err := c.authjwt.GenerateUserToken(&claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrServer)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, newWrap(authSignInOut{
		Token: token,
	}))
}
