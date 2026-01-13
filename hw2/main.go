package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	

	e.GET("/echo", func(c echo.Context) error {
		message := c.QueryParam("message")
		if message == "" {
			return c.String(http.StatusBadRequest, "Missing parameter")
		}
		return c.String(http.StatusOK, message)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
