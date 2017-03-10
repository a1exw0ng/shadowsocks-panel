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

	regCtx := controller.NewRegisterCtx()
	e.GET("/register", regCtx.Register)
	e.POST("/register", regCtx.PostRegister)
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
