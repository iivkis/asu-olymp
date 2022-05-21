package controllerV1

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	authjwt "github.com/iivkis/asu-olymp/internal/auth_jwt"
	"github.com/iivkis/asu-olymp/internal/repository"
)

/*default output struct*/
type DefaultOut struct {
	ID uint `json:"id"`
}

/*wrapper*/
type wrap struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func inWrap(data interface{}) *wrap {
	w := &wrap{Data: data}
	if _, ok := data.(*ControllerError); !ok {
		w.Status = true
	}
	return w
}

/*getters from ctx*/
func getUserClaims(ctx *gin.Context) (*authjwt.UserClaims, bool) {
	val, ok := ctx.Get("user_claims")
	return val.(*authjwt.UserClaims), ok
}

func getPayload(ctx *gin.Context) *repository.Payload {
	return ctx.MustGet("payload").(*repository.Payload)
}

/*validator*/
// key - field name, val - validate function
type validatorRules map[string]func(val interface{}) bool

func validator(m map[string]interface{}, v validatorRules) error {
	for key, fn := range v {
		if val, ok := m[key]; ok {
			if !reflect.ValueOf(val).IsNil() {
				if !fn(val) {
					return fmt.Errorf("invalid value in field `%s`", key)
				}
			} else {
				delete(m, key)
			}
		}
	}
	if len(m) == 0 {
		return fmt.Errorf("the number of updated fields is zero")
	}
	return nil
}
