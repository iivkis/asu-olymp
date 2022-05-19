package controllerV1

type ControllerError struct {
	Code        uint16 `json:"code"`
	Description string `json:"description"`
}

func newControllerError(code uint16, description string) *ControllerError {
	return &ControllerError{
		Code:        code,
		Description: description,
	}
}

func (e *ControllerError) Add(s string) *ControllerError {
	return &ControllerError{
		Code:        e.Code,
		Description: e.Description + ": " + s,
	}
}

/* ALL CONTROLLER ERRORS*/

var ErrServer = newControllerError(0, "server error")

var (
	ErrRecordNotFound = newControllerError(1, "record not found")
	ErrIncorrectData  = newControllerError(2, "incorrect data")
	ErrEmailRegistred = newControllerError(3, "email already registred")
)
