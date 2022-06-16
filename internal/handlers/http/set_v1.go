package httphr

import (
	"github.com/gin-gonic/gin"
	ctrlv1 "github.com/iivkis/asu-olymp/internal/controllers/v1"
)

func (h *HandlerHttp) setControllersV1(router *gin.RouterGroup, c *ctrlv1.Controller) {
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

			router.GET("", c.Middleware.ByApiKey(false), c.Tasks.Get)
			router.GET("/:id", c.Middleware.ByApiKey(false), c.Tasks.GetByID)

			router.POST("", c.Middleware.ByApiKey(true), c.Tasks.Post)
			router.PUT("/:id", c.Middleware.ByApiKey(true), c.Tasks.Put)
		}

		//QUESTIONS
		{
			router := router.Group("questions")

			router.GET("", c.Middleware.ByApiKey(false), c.Questions.Get)
			router.GET("/:id", c.Middleware.ByApiKey(false), c.Questions.GetByID)

			router.POST("", c.Middleware.ByApiKey(true), c.Questions.Post)
			router.PUT("/:id", c.Middleware.ByApiKey(true), c.Questions.Put)
		}

		//ANSWERS
		{
			router := router.Group("answers")

			router.GET("", c.Middleware.ByApiKey(true), c.Answers.Get)
			router.GET("/:id", c.Middleware.ByApiKey(true), c.Answers.GetByID)

			router.POST("", c.Middleware.ByApiKey(true), c.Answers.Post)
			router.PUT("/:id", c.Middleware.ByApiKey(true), c.Answers.Put)
		}

		{
			router := router.Group("check")
			router.POST("", c.Middleware.ByApiKey(false), c.CheckingAnswers.Post)

		}
	}
}
