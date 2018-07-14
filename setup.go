package main

import (
	"github.com/hack-boozer/boozer-api/db"
	"github.com/hack-boozer/boozer-api/db/migrations"
	mid "github.com/hack-boozer/boozer-api/middleware"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var database *gorm.DB

func setup(e *echo.Echo) {
	// DB接続
	database = db.ConnectDB()
	// migration
	migrations.Execute()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(mid.CORSMiddleware()))
}
