package handlerHTTP

import (
	"github.com/gin-gonic/gin"
	controllerV1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
)

func (h *HandlerHttp) setControllersV1(router *gin.RouterGroup, c *controllerV1.ControllerV1) {
	//AUTH
	{
		router.POST("/signUp", c.Auth.SignUp)
		router.POST("/signIn", c.Auth.SignIn)
	}

	//TASKS
	{
		router := router.Group("tasks")
		router.GET("/", c.Middleware.WithBearer(false), c.Tasks.Get)
	}
}
