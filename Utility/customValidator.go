package Utility

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cValidator *CustomValidator) Validate(i interface{}) error {

	if err := cValidator.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return nil
}
