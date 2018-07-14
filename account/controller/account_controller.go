package controller

import (
	"net/http"

	"github.com/hack-boozer/boozer-api/account/usecase"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// AccountController account controller
type AccountController struct {
	accountUse usecase.AccountUsecase
}

// NewAccountController mount account usecase
func NewAccountController(e *echo.Echo, accountUse usecase.AccountUsecase) {
	handler := &AccountController{
		accountUse: accountUse,
	}

	e.GET("/accounts/:id", handler.Get)
}

// Get get account
func (c *AccountController) Get(ctx echo.Context) error {
	id := uuid.FromStringOrNil(ctx.Param("id"))
	res, err := c.accountUse.GetAccount(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}
