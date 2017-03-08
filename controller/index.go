package controller

import (
	"github.com/labstack/echo"
	"net/http"
)

type PublicCtx struct {

}

func NewPublicCtx() *PublicCtx {
	return &PublicCtx{}
}

func (pub *PublicCtx) Index(c echo.Context) error {
	return c.String(http.StatusOK, "This is index page")
}
