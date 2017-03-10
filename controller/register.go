package controller

import (
	"github.com/labstack/echo"
	"fmt"
	"../models"
	"net/http"
	s "../session"
	"../auth"
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

func (pub *Register) PostRegister(c *s.Context) error {
	fmt.Println()
	fmt.Println(c.FormValue("password"))
	fmt.Println(c.FormValue("email"))

	username := c.FormValue("username")
	password := c.FormValue("password")

	if (false == models.CheckUserExistence(username)){
		u := models.GetUserByNicknamePwd(username, password)

		if u != nil{
			session := c.Session()
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, "/")
			return nil
		} else {
			c.Redirect(http.StatusMovedPermanently, "/register")
			return nil
		}
	}


	return nil
}