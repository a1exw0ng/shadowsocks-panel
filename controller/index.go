package controller

import (
	"github.com/labstack/echo"
	s "../session"
)

type PublicCtx struct {

}

func NewPublicCtx() *PublicCtx {
	return &PublicCtx{}
}

func (pub *PublicCtx) Index(c *s.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "index")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) Download(c *s.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "download_org")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) UOrder(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "uorder")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) Invite(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "invite")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) How(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "how")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}
func (pub *PublicCtx) User(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "user")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) Node(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "node")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) Buy(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "buy")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) ChangePW(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "changepw")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}


func (pub *PublicCtx) Tos(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "tos")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}

func (pub *PublicCtx) Logout(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "logout")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}