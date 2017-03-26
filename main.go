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
	e.Debug = true
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())
	e.Static("/static", "static")

	e.Use(session.Session())
	e.Use(auth.New(models.GenerateAnonymousUser))

	// Public pages
	pubCtx := controller.NewPublicCtx()

	e.GET("/", handler(pubCtx.Index))
	e.GET("/index.html", handler(pubCtx.Index))
	e.GET("/download_org.html", handler(pubCtx.Download))

	regCtx := controller.NewRegisterCtx()
	e.GET("/register", regCtx.Register)
	e.POST("/register", handler(regCtx.PostRegister))

	loginCtx := controller.NewLoginCtx()
	e.GET("/login", handler(loginCtx.Login))
	e.POST("/login", handler(loginCtx.PostLogin))

	e.GET("/logout", handler(loginCtx.LogoutHandler))
	e.GET("/logout.html", handler(loginCtx.LogoutHandler))

	// Admin pages
	adminCtx := controller.NewAdminCtx()
	adminGroup := e.Group("/admin")
	adminGroup.Use(auth.LoginRequired())
	{
		adminGroup.GET("/index", adminCtx.Index)
	}

	// User pages
	userCtx := controller.NewUserCtx()

	userGroup := e.Group("/user")
	userGroup.Use(auth.LoginRequired())
	{
		//userGroup.GET("/index", userCtx.Index)
		userGroup.GET("/download.html", userCtx.Download)
		userGroup.GET("/uorder.html", userCtx.UOrder)
		userGroup.GET("/invite.html", userCtx.Invite)
		userGroup.GET("/how.html", userCtx.How)
		userGroup.GET("/user.html", handler(userCtx.User))
		userGroup.GET("/buy.html", userCtx.Buy)
		userGroup.GET("/changepw.html", userCtx.ChangePW)
		userGroup.GET("/node.html", userCtx.Node)
		userGroup.GET("/tos.html", userCtx.Tos)
		userGroup.GET("/logout", handler(userCtx.LogoutHandler))
		userGroup.GET("/logout.html", handler(userCtx.LogoutHandler))
	}

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
