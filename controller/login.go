package controller

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	s "../session"
	"../auth"
	"../models"
)

func Login(c echo.Context) error {
	c.Set("tmpl", "login")
	c.Set("data", map[string]interface{}{
		"title":"Login",
	})
	return nil
}

func PostLogin(c *s.Context) error {
	fmt.Println(c.FormValue("username"))
	fmt.Println(c.FormValue("password"))
	fmt.Println(c.FormValue("email"))

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	a := c.Auth()
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	u := models.GetUserByNicknamePwd("hello", "password")
	if u != nil{
		session := c.Session()
		err := auth.AuthenticateSession(session, u)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.Redirect(http.StatusMovedPermanently, "/")
		return nil
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return nil
	}

}

