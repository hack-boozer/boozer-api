package main

import (
	"fmt"
	"os"

	accountC "github.com/hack-boozer/boozer-api/account/controller"
	accountR "github.com/hack-boozer/boozer-api/account/repository"
	accountU "github.com/hack-boozer/boozer-api/account/usecase"
	hello "github.com/hack-boozer/boozer-api/hello/controller"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	setup(e)

	accountRepo := accountR.NewAccountRepository(database)

	accountUse := accountU.NewAccountUsecase(accountRepo)

	hello.NewHelloController(e)
	accountC.NewAccountController(e, accountUse)

	fmt.Println("vim-go")
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
