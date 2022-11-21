package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ShirtsPost(c echo.Context) error {
	return c.String(http.StatusCreated, "ShirtCreated")
}
