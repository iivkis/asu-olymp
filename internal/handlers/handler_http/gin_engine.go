package handlerHTTP

import "github.com/gin-gonic/gin"

func newGinEngine() (engine *gin.Engine) {
	engine = gin.Default()
	return
}
