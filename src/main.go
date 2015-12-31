package main

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

// Handler
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	initCofig()

	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/", hello)

	e.Run(viper.GetString("port"))
}
