package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ShirtPostHandler(c echo.Context) error {
	// name := c.FormValue("name")
	// email := c.FormValue("email")

	return c.JSON(http.StatusCreated, "ShirtCreated")
}
