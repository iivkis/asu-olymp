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

type AuthSignUpBody struct {
	Email    string `json:"email" binding:"required,min=3,max=100" minLength:"3" maxLength:"100" example:"example@mail.ru"`
	Password string `json:"password" binding:"required,min=4,max=50" minLength:"4" maxLength:"50" example:"qwerty27"`
	FullName string `json:"full_name" binding:"required,min=1,max=100" minLength:"1" maxLength:"100" example:"Фёдоров И.С."`
}

//@Summary Create a new user profile
//@Tags auth
//@ID SignUp
//@Param body body AuthSignUpBody true "sign up data"
//@Success 201 {object} wrap{data=DefaultOut}
//@Failure 400
//@Failure 500
//@Router /signUp [post]
func (c *AuthController) SignUp(ctx *gin.Context) {
	var body AuthSignUpBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	if c.repository.Users.Exists(&repository.UserModel{Email: body.Email}) {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrEmailRegistred))
		return
	}

	model := repository.UserModel{
		Email:    body.Email,
		FullName: body.FullName,
		Password: body.Password,
	}

	if err := c.repository.Users.Cursor().Create(&model).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, inWrap(ErrServer.Add(err.Error())))
		return
	}

	ctx.JSON(http.StatusCreated, inWrap(DefaultOut{ID: model.ID}))
}

type AuthSignInBody struct {
	Email    string `json:"email" binding:"required,min=3,max=100" minLength:"3" maxLength:"100" example:"example@mail.ru"`
	Password string `json:"password" binding:"required,min=4,max=50" minLength:"4" maxLength:"50" example:"qwerty27"`
}

type AuthSignInOut struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsImlzcyI6ImFzdS1vbHltcCJ9.NPFZIvICrpfdqUlbr5vfvRMCHgbKj28eXmLjftWPjyc"`
}

//@Summary Sign in user profile
//@Tags auth
//@ID SignIn
//@Param body body AuthSignInBody true "sign in data"
//@Success 200 {object} wrap{data=AuthSignInOut}
//@Failure 400
//@Failure 500
//@Router /signIn [post]
func (c *AuthController) SignIn(ctx *gin.Context) {
	var body AuthSignInBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	user, err := c.repository.Users.SignUpByEmail(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData))
		return
	}

	claims := authjwt.UserClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			Issuer: "asu-olymp",
		},
	}

	token, err := c.authjwt.GenerateUserToken(&claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrServer)
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, inWrap(AuthSignInOut{
		Token: token,
	}))
}
