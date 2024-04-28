package api

import (
    "fmt"
    "net/http"
	"github.com/gorilla/sessions"
    // "time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
)

func Register(c echo.Context) error {
    return c.String(http.StatusOK, "Signed up")
}

type LoginBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResp struct {
    Username string `json:"username"`
}

func LogIn(c echo.Context) error {
    var loginBody LoginBody
    if err := c.Bind(&loginBody); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to parse params: %s", err))
    }
    fmt.Printf("username: %s pass: %s\n", loginBody.Username, loginBody.Password)

    sess, err := session.Get("session", c)
    if err != nil {
        return err
    }
    fmt.Printf("session id was generated and is %s\n", sess.ID)

    sess.Options = &sessions.Options{
        Path: "/",
        MaxAge: 86400 * 7,
        HttpOnly: true,
    }

    sess.Values["userID"] = getUserID()

    if err := sess.Save(c.Request(), c.Response()); err != nil {
			return err
	}

    fmt.Println("SessionID after saving is", sess.ID)

    return c.JSON(http.StatusOK, LoginResp{
        Username: loginBody.Username,
    })
}

func getUserID() int {
    return 1
}

func LogOut(c echo.Context) error {
    return c.String(http.StatusOK, "Logged out")
}

