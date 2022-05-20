package controllerV1

import (
	"github.com/gin-gonic/gin"
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
)

func getUserClaims(ctx *gin.Context) (*authjwt.UserClaims, bool) {
	val, ok := ctx.Get("user_claims")
	return val.(*authjwt.UserClaims), ok
}
