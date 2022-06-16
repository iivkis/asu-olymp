package ctrlv1

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/iivkis/asu-olymp/internal/auth"
	"github.com/iivkis/asu-olymp/internal/repository"
)

/*default output struct*/
//@Description Record ID
type DefaultOut struct {
	ID uint `json:"id" minimum:"0"`
}

/*wrapper*/
//@Description Standard wrapper for responses
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
func getUserClaims(ctx *gin.Context) (*auth.UserClaims, bool) {
	val, ok := ctx.Get("user_claims")
	return val.(*auth.UserClaims), ok
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
