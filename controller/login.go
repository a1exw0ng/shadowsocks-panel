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

	fmt.Println(c.FormParams())
	//name := c.FormValue("username")
	password := c.FormValue("passwd")
	email := c.FormValue("email")
	redirect := c.QueryParam(auth.RedirectParam)
	fmt.Println(c.QueryString())
	fmt.Println("auth.RedirectParam", auth.RedirectParam, redirect)
	if redirect == "" {
		redirect = "/"
	}

	a := c.Auth()
	fmt.Println("redirect--", redirect)
	if a.User.IsAuthenticated() {
		fmt.Println("is authenticated")
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	//u := models.GetUserByNicknamePwd(name, password)
	u := models.GetUserByMailPwd(email, password)
	fmt.Println("user=", u)
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

