package api

import (
    "fmt"
    "net/http"
	// "github.com/gorilla/sessions"
    // "time"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo-contrib/session"
)

func Register(c echo.Context) error {
    return c.String(http.StatusOK, "Signed up")
}

type Message struct {
    Msg string `json:"msg"`
}

func LogIn(c echo.Context) error {
    fmt.Println("logging in")
    return c.JSON(http.StatusOK, Message{Msg: "logged in"})
}

func LogOut(c echo.Context) error {
    return c.String(http.StatusOK, "Logged out")
}

