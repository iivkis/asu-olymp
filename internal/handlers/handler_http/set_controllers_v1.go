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
		router.POST("/cookie", c.Auth.FromCookie)
		router.DELETE("/cookie", c.Auth.DeleteCookie)
	}

	//route with tasks methods
	{
		router := router.Group("t")
		router.Use(c.Middleware.Payload)

		//TASKS
		{
			router := router.Group("tasks")

			router.GET("", c.Middleware.WithApiKey(false), c.Tasks.Get)
			router.GET("/:id", c.Middleware.WithApiKey(false), c.Tasks.GetByID)

			router.POST("", c.Middleware.WithApiKey(true), c.Tasks.Post)
			router.PUT("/:id", c.Middleware.WithApiKey(true), c.Tasks.Put)
		}

		//QUESTIONS
		{
			router := router.Group("questions")

			router.GET("", c.Middleware.WithApiKey(false), c.Questions.Get)
			router.GET("/:id", c.Middleware.WithApiKey(false), c.Questions.GetByID)

			router.POST("", c.Middleware.WithApiKey(true), c.Questions.Post)
			router.PUT("/:id", c.Middleware.WithApiKey(true), c.Questions.Put)
		}

		//ANSWERS
		{
			router := router.Group("answers")

			router.GET("", c.Middleware.WithApiKey(true), c.Answers.Get)
			router.GET("/:id", c.Middleware.WithApiKey(true), c.Answers.GetByID)

			router.POST("", c.Middleware.WithApiKey(true), c.Answers.Post)
			router.PUT("/:id", c.Middleware.WithApiKey(true), c.Answers.Put)
		}

		{
			router := router.Group("check")
			router.POST("", c.Middleware.WithApiKey(false), c.CheckingAnswers.Post)

		}
	}
}
