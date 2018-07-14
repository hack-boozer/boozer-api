package main

import (
	"fmt"
	"os"

	hello "github.com/hack-boozer/boozer-api/hello/controller"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	setup(e)

	hello.NewHelloController(e)

	fmt.Println("vim-go")
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
