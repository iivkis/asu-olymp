package httphr

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/docs"
	ctrlv1 "github.com/iivkis/asu-olymp/internal/controllers/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Config struct {
	ControllerV1 *ctrlv1.Controller
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
	//use corse
	h.withCors()

	//set api v1
	h.setControllersV1(h.engine.Group("/api/v1"), h.cfg.ControllerV1)

	//set swagger
	docs.SwaggerInfo.Schemes = []string{"http"}
	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (h *HandlerHttp) withCors() {
	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "Authorization")
	cfg.AllowOrigins = append(cfg.AllowOrigins, "http://localhost")
	cfg.AllowCredentials = true

	h.engine.Use(cors.New(cfg))
}
