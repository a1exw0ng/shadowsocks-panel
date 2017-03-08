package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type AdminCtx struct {

}

func NewAdminCtx() *AdminCtx{
	ac := AdminCtx{}
	return &ac
}

func (ac *AdminCtx) Index(c echo.Context) error {
	return c.String(http.StatusOK, "You're admin")
}