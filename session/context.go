package session

import "github.com/labstack/echo"
import "github.com/hobo-go/echo-mw/session"
import (
	"../auth"
	"fmt"
)

func NewContext() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("NewContext Middleware ")
			ctx := &Context{c}
			return h(ctx)
		}
	}
}

type Context struct {
	echo.Context
}

func (ctx *Context) Session() session.Session {
	return session.Default(ctx)
}

func (ctx *Context) Auth() auth.Auth {
	return auth.Default(ctx)
}

