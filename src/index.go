package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/malls/thankyou-api/printful"
	"github.com/malls/thankyou-api/src/controllers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/shirt", controllers.ShirtPostHandler)
	e.GET("/printful-categories", func(c echo.Context) error {
		data := printful.GetVariantIds()

		return c.JSONPretty(http.StatusOK, data, "  ")
	})
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "Running Okay")
	})

	e.Logger.Fatal(e.Start(":1326"))
}
