package controller

import (
	"github.com/labstack/echo"
)

type AdminCtx struct {

}

func NewAdminCtx() *AdminCtx{
	ac := AdminCtx{}
	return &ac
}

func (ac *AdminCtx) Index(c echo.Context) error {
	c.Set("tmpl", "admin")
	c.Set("data", map[string]interface{}{
		"hello":"World",
	})
	return nil
}