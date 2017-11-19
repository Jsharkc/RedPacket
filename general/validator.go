package general

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// EchoValidator - validator
type EchoValidator struct {
	validator *validator.Validate
}

// Validate - validate input param function
func (ev *EchoValidator) Validate(i interface{}) error {
	return ev.validator.Struct(i)
}

// NewEchoValidator - create new Validator
func NewEchoValidator() echo.Validator {
	return &EchoValidator{
		validator: validator.New(),
	}
}
