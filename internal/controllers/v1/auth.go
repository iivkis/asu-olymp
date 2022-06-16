package ctrlv1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/iivkis/asu-olymp/internal/auth"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type AuthController struct {
	repository *repository.Repository
	authjwt    *auth.AuthorizationJWT
}

func NewAuthController(repository *repository.Repository, auth *auth.AuthorizationJWT) *AuthController {
	return &AuthController{repository: repository, authjwt: auth}
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

	user, err := c.repository.Users.SignInByEmail(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData))
		return
	}

	claims := auth.UserClaims{
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

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "apikey",
		Value:    token,
		Path:     "/",
		MaxAge:   24 * 60 * 60,
		HttpOnly: true,
	})

	ctx.JSON(http.StatusOK, inWrap(AuthSignInOut{
		Token: token,
	}))
}

type AuthFromCookieOut struct {
	Token string `json:"token"`
}

//@Summary Get ApiKey by cookie
//@Tags auth
//@ID FromCookie
//@Success 200 {object} wrap{data=AuthFromCookieOut}
//@Failure 400
//@Router /cookie [post]
func (c *AuthController) FromCookie(ctx *gin.Context) {
	apikey, err := ctx.Cookie("apikey")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrIncorrectData.Add("cookie undefined"))
		return
	}
	ctx.JSON(http.StatusOK, inWrap(AuthFromCookieOut{Token: apikey}))
}

//@Summary Delete cookie
//@Tags auth
//@ID DeleteCookie
//@Success 200
//@Router /cookie [delete]
func (c *AuthController) DeleteCookie(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "apikey",
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
}
