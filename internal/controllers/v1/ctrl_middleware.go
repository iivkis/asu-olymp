package ctrlv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/auth"
	"github.com/iivkis/asu-olymp/internal/repository"
)

type MiddlewareController struct {
	authjwt *auth.AuthorizationJWT
}

func NewMiddlewareController(authjwt *auth.AuthorizationJWT) *MiddlewareController {
	return &MiddlewareController{authjwt: authjwt}
}

func (c *MiddlewareController) ByApiKey(mandatory bool) func(ctx *gin.Context) {
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
	OffsetID uint `form:"offset_id" json:"offset_id" binding:"min=0"`
	Limit    int  `form:"limit" json:"limit" binding:"min=0,max=1000"`
}

func (c *MiddlewareController) Payload(ctx *gin.Context) {
	var query PayloadQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, inWrap(ErrIncorrectData.Add(err.Error())))
		return
	}

	if query.Limit == 0 {
		query.Limit = 1000
	}

	ctx.Set("payload", &repository.Payload{
		OffsetID: query.OffsetID,
		Limit:    query.Limit,
	})
}
