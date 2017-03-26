package session

import (
	"github.com/labstack/echo"

	es "github.com/hobo-go/echo-mw/session"

	"fmt"
)

func Session() echo.MiddlewareFunc {
	//store := es.NewCookieStore([]byte("secret"))
	//return es.New("mysession", store)
	fmt.Println("Session() Middleware")
	store := es.NewFilesystemStore("testxxx", []byte("secret-key"))
	return es.New("mysession", store)
}
