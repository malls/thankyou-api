package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/malls/thankyou-api/src/controllers"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/shirt", controllers.ShirtPostHandler)
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Running Okay")
	})

	e.Logger.Fatal(e.Start(":1326"))
}
