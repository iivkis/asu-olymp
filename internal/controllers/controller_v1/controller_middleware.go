package controllerV1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type MiddlewareController struct {
	authjwt *authjwt.AuthJWT
}

func NewMiddlewareController(authjwt *authjwt.AuthJWT) *MiddlewareController {
	return &MiddlewareController{authjwt: authjwt}
}

func (c *MiddlewareController) Bearer(mandatory bool) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "" {
			claims, err := c.authjwt.ParseUserToken(token)
			if err != nil {
				ctx.AbortWithStatusJSON(401, inWrap(ErrIncorrectData.Add(err.Error())))
				return
			}
			ctx.Set("user_claims", claims)
		} else if mandatory {
			ctx.AbortWithStatus(401)
			return
		}
	}
}

type PayloadQuery struct {
	OffsetID uint `form:"offset_id" binding:"min=0"`
	Limit    int  `form:"limit" binding:"min=0,max=1000"`
}

func (c *MiddlewareController) Payload(ctx *gin.Context) {
	var query PayloadQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	ctx.Set("payload", &repository.Payload{
		OffsetID: query.OffsetID,
		Limit:    query.Limit,
	})
}
