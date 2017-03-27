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

	username := c.FormValue("email") //默认用户名和密码一致
	password := c.FormValue("passwd")
	email := c.FormValue("email")

	if (false == models.CheckUserExistence(username)){

		err := models.CreateUserByNamePwdEmail(username, password, email)
		if err != nil{
			fmt.Println("Insert into database failed", err.Error())
			return err
		}
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