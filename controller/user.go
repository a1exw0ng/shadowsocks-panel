package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type UserCtx struct {

}

func NewUserCtx() *UserCtx{
	ac := AdminCtx{}
	return &ac
}

func (ac *UserCtx) Index(c echo.Context) error {
	return c.String(http.StatusOK, "You're a user")
}