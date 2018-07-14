package main

import (
	"fmt"
	"os"

	accountC "github.com/hack-boozer/boozer-api/account/controller"
	accountR "github.com/hack-boozer/boozer-api/account/repository"
	accountU "github.com/hack-boozer/boozer-api/account/usecase"
	hello "github.com/hack-boozer/boozer-api/hello/controller"
	photoR "github.com/hack-boozer/boozer-api/photo/repository"
	postC "github.com/hack-boozer/boozer-api/post/controller"
	postR "github.com/hack-boozer/boozer-api/post/repository"
	postU "github.com/hack-boozer/boozer-api/post/usecase"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	setup(e)

	accountRepo := accountR.NewAccountRepository(database)
	postRepo := postR.NewPostRepository(database)
	photoRepo := photoR.NewPhotoRepository(database)

	accountUse := accountU.NewAccountUsecase(accountRepo)
	postUse := postU.NewPostUsecase(postRepo, photoRepo)

	hello.NewHelloController(e)
	accountC.NewAccountController(e, accountUse)
	postC.NewPostController(e, postUse)

	fmt.Println("vim-go")
	port := ":" + os.Getenv("PORT")
	e.Logger.Fatal(e.Start(port))
}
