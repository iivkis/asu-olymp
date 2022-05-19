package controllerV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ping struct{}

func newPing() *ping {
	return &ping{}
}

func (c *ping) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, newWrap("pong"))
}

func (c *ping) GetErrPong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, newWrap(ErrServer.Add("ok!")))
}
