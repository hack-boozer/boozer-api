package controller

import (
	"net/http"

	"github.com/hack-boozer/boozer-api/post"
	"github.com/hack-boozer/boozer-api/post/usecase"
	"github.com/labstack/echo"
)

// PostController post controller
type PostController struct {
	postUse usecase.PostUsecase
}

// NewPostController mount post controller
func NewPostController(e *echo.Echo, post usecase.PostUsecase) {
	handler := &PostController{
		postUse: post,
	}

	e.POST("/posts", handler.Create)
}

// Create create post
func (c *PostController) Create(ctx echo.Context) error {
	req := post.Create{}
	err := ctx.Bind(&req)
	res, err := c.postUse.Create(&req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}
