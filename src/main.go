package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {

		return c.HTML(http.StatusOK, "<h1>Works fine</h2>")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK,
			struct {
				Status  string
				Mensage string
			}{Status: "OK", Mensage: "Works Fine"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8081"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
