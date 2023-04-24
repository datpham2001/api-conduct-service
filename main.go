package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello world")
	})

	e.Logger.Fatal(e.Start(":8080"))

}
