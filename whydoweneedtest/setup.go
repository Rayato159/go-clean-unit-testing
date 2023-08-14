package whydoweneedtest

import (
	"net/http/httptest"

	"github.com/Rayato159/go-clean-unit-testing/pkg/utils"
	"github.com/labstack/echo/v4"
)

func NewEchoContext[T any](method, endpoint string, body T) echo.Context {
	e := echo.New()

	req := httptest.NewRequest(method, endpoint, utils.ConvertObjToStringReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	return e.NewContext(req, httptest.NewRecorder())
}
