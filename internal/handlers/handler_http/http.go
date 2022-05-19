package handlerHTTP

import (
	"github.com/gin-gonic/gin"
	cv1 "github.com/iivkis/asu-olymp/internal/controllers/controller_v1"
)

type Config struct {
	ControllerV1 *cv1.ControllerV1
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
	h.setControllersV1(h.engine.Group("/api/v1"), h.cfg.ControllerV1)
}
