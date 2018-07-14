package main

import (
	"git.si-media.biz/hack-boozer/boozer-api/db/migrations"
	"github.com/hack-boozer/boozer-api/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setup(e *echo.Echo) {
	// DB接続
	database = db.ConnectDB()
	// migration
	migrations.Execute()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
