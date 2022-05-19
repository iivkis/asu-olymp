package handlerHTTP

import (
	"github.com/gin-gonic/gin"
	cv1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
)

func (h *HandlerHttp) setControllersV1(router *gin.RouterGroup, controller *cv1.ControllerV1) {
	//AUTH
	{
		router.POST("/signUp", controller.Auth.SignUp)
		router.POST("/signIn", controller.Auth.SignIn)
	}
}
