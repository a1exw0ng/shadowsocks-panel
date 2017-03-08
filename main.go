package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"./controller"
	"./render"
)

func main()  {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:"time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Renderer = render.LoadTemplates()
	e.Static("/static", "static")

	// Admin pages
	adminCtx := controller.NewAdminCtx()
	adminGroup := e.Group("/admin")
	adminGroup.GET("/index", adminCtx.Index)

	// User pages
	userCtx := controller.NewUserCtx()

	userGroup := e.Group("/user")
	userGroup.GET("/index", userCtx.Index)

	// Public pages
	pubCtx := controller.NewPublicCtx()

	e.GET("/", pubCtx.Index)

	e.Logger.Fatal(e.Start(":8082"))
}