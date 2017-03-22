package controller

import (
	"github.com/labstack/echo"
	s "../session"
	"fmt"
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

func (pub *UserCtx) Download(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "download")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}


func (pub *UserCtx) UOrder(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "uorder")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *UserCtx) Invite(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "invite")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *UserCtx) How(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "how")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}
func (pub *UserCtx) User(c *s.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	a := c.Auth()
	fmt.Println("userId=", a.User.UniqueId())

	c.Set("tmpl", "user")
	c.Set("data", map[string]interface{}{
		"Email":"Email",
	})
	return nil
}

func (pub *UserCtx) Node(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "node")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *UserCtx) Buy(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "buy")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *UserCtx) ChangePW(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "changepw")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}


func (pub *UserCtx) Tos(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "tos")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *UserCtx) Logout(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "logout")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})

	return nil
}