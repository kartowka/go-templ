package main

import (
	"github.com/antfley/GOT/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/context"
)

func main() {
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	}))
	app.Use(withUser)
	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "antfley@gmail.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
