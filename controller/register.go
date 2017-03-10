package controller

import (
	"github.com/labstack/echo"
	"fmt"
)

type Register struct {

}

func NewRegisterCtx() *Register {
	return &Register{}
}

func (pub *Register) Register(c echo.Context) error {

	c.Set("tmpl", "register")
	c.Set("data", map[string]interface{}{
		"title":"Register",
	})
	return nil
}

func (pub *Register) PostRegister(c echo.Context) error {
	fmt.Println(c.FormValue("username"))
	fmt.Println(c.FormValue("password"))
	fmt.Println(c.FormValue("email"))
	return nil
}