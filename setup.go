package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setup(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
