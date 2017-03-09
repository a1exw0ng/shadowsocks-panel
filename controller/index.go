package controller

import (
	"github.com/labstack/echo"
)

type PublicCtx struct {

}

func NewPublicCtx() *PublicCtx {
	return &PublicCtx{}
}

func (pub *PublicCtx) Index(c echo.Context) error {

	//return c.String(http.StatusOK, "This is index page")
	c.Set("tmpl", "index")
	c.Set("data", map[string]interface{}{
		"title":"About",
	})
	return nil
}
