package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type PostShirtSuccess struct {
	Message string `json:"message" xml:"message"`
	Url     string `json:"url" xml:"url"`
}

func ShirtPostHandler(c echo.Context) error {

	image, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Source
	src, err := image.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./images/" + image.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	m := &PostShirtSuccess{
		Message: "Success",
		Url:     image.Filename,
	}

	return c.JSONPretty(http.StatusCreated, m, "  ")

}
