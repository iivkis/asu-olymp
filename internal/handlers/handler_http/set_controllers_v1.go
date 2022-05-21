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
		router.Use(c.Middleware.Payload)

		router.GET("", c.Middleware.Bearer(false), c.Tasks.Get)
		router.GET("/:id", c.Middleware.Bearer(false), c.Tasks.GetByID)
		router.POST("", c.Middleware.Bearer(true), c.Tasks.Post)
		router.PUT("/:id", c.Middleware.Bearer(true), c.Tasks.Put)
	}

	//QUESTIONS
	{
		router := router.Group("questions")
		router.Use(c.Middleware.Payload)

		router.GET("", c.Middleware.Bearer(false), c.Questions.Get)
		router.POST("", c.Middleware.Bearer(true), c.Questions.Post)
	}

	//ANSWERS
	{
		router := router.Group("answers")
		router.Use(c.Middleware.Payload)

		router.GET("", c.Middleware.Bearer(true), c.Answers.Get)
		router.POST("", c.Middleware.Bearer(true), c.Answers.Post)
	}
}
