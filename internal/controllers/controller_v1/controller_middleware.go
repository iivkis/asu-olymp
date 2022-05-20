package controllerV1

import (
	"github.com/gin-gonic/gin"
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
)

type MiddlewareController struct {
	authjwt *authjwt.AuthJWT
}

func NewMiddlewareController(authjwt *authjwt.AuthJWT) *MiddlewareController {
	return &MiddlewareController{authjwt: authjwt}
}

func (c *MiddlewareController) WithBearer(mandatory bool) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "" {
			claims, err := c.authjwt.ParseUserToken(token)
			if err != nil {
				ctx.AbortWithStatusJSON(401, newWrap(ErrIncorrectData.Add(err.Error())))
				return
			}
			ctx.Set("user_claims", claims)
		} else if mandatory {
			ctx.AbortWithStatus(401)
			return
		}
	}
}
