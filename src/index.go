package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/malls/thankyou-api/controllers"
)

func main() {

	controllers.ShirtsPost()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
