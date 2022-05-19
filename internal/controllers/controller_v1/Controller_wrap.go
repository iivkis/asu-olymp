package controllerV1

type wrap struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func newWrap(data interface{}) *wrap {
	w := &wrap{Data: data}
	if _, ok := data.(*ControllerError); !ok {
		w.Status = true
	}
	return w
}
