package handlerHTTP

import (
	"github.com/gin-gonic/gin"
	cv1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
)

func (h *HandlerHttp) setControllersV1(router *gin.RouterGroup, c *cv1.ControllerV1) {
	{
		router.GET("/ping", c.Ping.Get)
		router.GET("/ping-pong", c.Ping.GetErrPong)
	}

	//AUTH
	{
		router.POST("/signUp", c.Auth.SignUp)
	}
}
