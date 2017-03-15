package controller

import (
	"github.com/labstack/echo"
)

type UserCtx struct {

}

func NewUserCtx() *UserCtx{
	ac := &UserCtx{}
	return ac
}

func (ac *UserCtx) Index(c echo.Context) error {
	c.Set("tmpl", "user")
	c.Set("data", map[string]interface{}{
		"Password":"Helloworld",
		"Port":900,
	})
	return nil
}