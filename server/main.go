package main

import (
	"net/http"

	"github.com/flippinberger/djmocracy/api"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:5173", "http://localhost"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

    // TODO update secret here with env var secret
    e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Nuxt App!")
	})

    e.POST("/signup", api.Register)
    e.POST("/login", api.LogIn)
    e.POST("/logout", api.LogOut)

	e.Logger.Fatal(e.Start(":1323"))
}
