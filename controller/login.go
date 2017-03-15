package controller

import (
	"github.com/labstack/echo"
	"net/http"
	s "../session"
	"../auth"
	"../models"
	"fmt"
)

func Login(c echo.Context) error {
	c.Set("tmpl", "login")
	c.Set("data", map[string]interface{}{
		"title":"Login",
	})
	return nil
}

func PostLogin(c *s.Context) error {

	name := c.FormValue("username")
	password := c.FormValue("password")
	//email := c.FormValue("email")
	redirect := c.QueryParam(auth.RedirectParam)
	fmt.Println(c.QueryString())
	fmt.Println("auth.RedirectParam", auth.RedirectParam, redirect)
	if redirect == "" {
		redirect = "/"
	}

	a := c.Auth()
	fmt.Println("redirect--", redirect)
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	u := models.GetUserByNicknamePwd(name, password)
	if u != nil{
		session := c.Session()
		err := auth.AuthenticateSession(session, u)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		fmt.Println("redirect=", redirect)
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return nil
	}

}

