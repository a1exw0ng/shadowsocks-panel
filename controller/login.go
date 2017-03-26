package controller

import (
	"net/http"
	s "../session"
	"../auth"
	"../models"
	"fmt"
)

type Response struct {
	Ok bool `json:ok`
	Msg string `json:msg`
}

type Login struct {

}

func NewLoginCtx() *Login {
	return &Login{}
}
func (l *Login)Login(c *s.Context) error {
	fmt.Println("$$$$Login.go - Login")

	//c.Set("tmpl", "login")
	//c.Set("data", map[string]interface{}{
	//	"title":"Login",
	//})
	//return nil
	redirect := c.QueryParam(auth.RedirectParam)

	a := c.Auth()
	if a.User.IsAuthenticated() {
		fmt.Println("###Login.go - Login=", a.User.IsAuthenticated())
		if redirect == "" {
			redirect = "/user/user.html"
		}
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	c.Set("tmpl", "login")
	c.Set("data", map[string]interface{}{
		"title":         "Login",
	})

	return nil
}

func (l *Login)PostLogin(c *s.Context) error {

	resp := Response{Ok:true, Msg:"登陆成功"}
	fmt.Println(c.FormParams())
	//name := c.FormValue("username")
	password := c.FormValue("passwd")
	email := c.FormValue("email")
	redirect := c.QueryParam(auth.RedirectParam)
	fmt.Println("auth.RedirectParam", auth.RedirectParam, redirect)
	if redirect == "" {
		redirect = "/"
	}

	a := c.Auth()

	if a.User.IsAuthenticated() {
		fmt.Println("PostLogin - is authenticated")
		c.JSON(http.StatusOK, resp)
		return nil
	}

	//u := models.GetUserByNicknamePwd(name, password)
	u := models.GetUserByMailPwd(email, password)
	fmt.Println("PostLogin - user=", u)
	if u != nil{
		session := c.Session()
		err := auth.AuthenticateSession(session, u)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, resp)
		return nil
	} else {
		fmt.Println("Login Failed Redirect to /login")
		c.Redirect(http.StatusOK, "/login")
		return nil
	}

}

func (l *Login)LogoutHandler(c *s.Context) error {
	fmt.Println("$$$ login.go-LogoutHandler=", c)
	session := c.Session()
	fmt.Println("$$$ login.go-LogoutHandler=", session)

	a := c.Auth()
	auth.Logout(session, a.User)

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}
