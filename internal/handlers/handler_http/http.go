package handlerHTTP

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/docs"
	controllerV1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Config struct {
	ControllerV1 *controllerV1.ControllerV1
}

type HandlerHttp struct {
	engine *gin.Engine
	cfg    *Config
}

func NewHandlerHTTP(cfg *Config) *HandlerHttp {
	h := &HandlerHttp{
		engine: newGinEngine(),
		cfg:    cfg,
	}
	h.init()
	return h
}

func (h *HandlerHttp) Engine() *gin.Engine {
	return h.engine
}

func (h *HandlerHttp) init() {
	h.engine.Use(cors.Default())

	//set api v1
	h.setControllersV1(h.engine.Group("/api/v1"), h.cfg.ControllerV1)

	//set swagger
	docs.SwaggerInfo.Schemes = []string{"http"}
	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
