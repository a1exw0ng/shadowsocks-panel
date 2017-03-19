package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./controller"
	"./render"
	"./auth"
	"./session"
	"./models"
)

func main()  {

	models.InitDB("test.db")

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:"time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Context自定义
	e.Use(session.NewContext())

	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())
	e.Static("/static", "static")

	e.Use(session.Session())
	e.Use(auth.New(models.GenerateAnonymousUser))

	// Admin pages
	adminCtx := controller.NewAdminCtx()
	adminGroup := e.Group("/admin")
	adminGroup.Use(auth.LoginRequired())
	adminGroup.GET("/index", adminCtx.Index)

	// User pages
	userCtx := controller.NewUserCtx()

	userGroup := e.Group("/user")
	userGroup.Use(auth.LoginRequired())
	userGroup.GET("/index", userCtx.Index)

	// Public pages
	pubCtx := controller.NewPublicCtx()

	e.GET("/", pubCtx.Index)
	e.GET("/index.html", pubCtx.Index)
	e.GET("/download.html", pubCtx.Download)
	e.GET("/uorder.html", pubCtx.UOrder)
	e.GET("/invite.html", pubCtx.Invite)
	e.GET("/how.html", pubCtx.How)
	e.GET("/user.html", pubCtx.User)
	e.GET("/buy.html", pubCtx.Buy)
	e.GET("/changepw.html", pubCtx.ChangePW)
	e.GET("/node.html", pubCtx.Node)
	e.GET("/tos.html", pubCtx.Tos)
	e.GET("/logout.html", pubCtx.Logout)

	regCtx := controller.NewRegisterCtx()
	e.GET("/register", regCtx.Register)
	e.POST("/register", handler(regCtx.PostRegister))
	e.GET("/login", controller.Login)
	e.POST("/login", handler(controller.PostLogin))

	e.Logger.Fatal(e.Start(":8082"))
}

type (
	HandlerFunc func(*session.Context) error
)

/**
 * 自定义Context的Handler
 */
func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*session.Context)
		return h(ctx)
	}
}
