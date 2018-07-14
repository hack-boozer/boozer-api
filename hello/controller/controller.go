package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloController hello world controller
type HelloController struct{}

// NewHelloController mount hello controller
func NewHelloController(e *echo.Echo) {
	handler := &HelloController{}

	e.GET("/hello", handler.Hello)
}

// Hello hello world
func Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello, World!")
}
