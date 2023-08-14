package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	contextWrapperUtils interface {
		Bind(data any) error
	}

	contextWrapper struct {
		Context   echo.Context
		validator *validator.Validate
	}
)

func ContextWrapper(c echo.Context) contextWrapperUtils {
	return &contextWrapper{
		Context:   c,
		validator: validator.New(),
	}
}

func (c *contextWrapper) Bind(data any) error {
	if err := c.Context.Bind(data); err != nil {
		return err
	}

	if err := c.validator.Struct(data); err != nil {
		return err
	}
	return nil
}
