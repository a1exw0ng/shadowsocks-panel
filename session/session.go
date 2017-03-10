package session

import (
	"github.com/labstack/echo"

	es "github.com/hobo-go/echo-mw/session"

)

func Session() echo.MiddlewareFunc {
	store := es.NewCookieStore([]byte("secret"))
	return es.New("mysession", store)
}
