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
    var loginBody LoginBody
    if err := c.Bind(&loginBody); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to parse params: %s", err))
    }
    fmt.Printf("username: %s pass: %s\n", loginBody.Username, loginBody.Password)

    //TODO create new user in db

    err := addSessionForUserId(c, fmt.Sprintf("%d", getUserID()))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Couldn't add session in register: %s", err))
    }

    return c.JSON(http.StatusOK, LoginResp{
        Username: loginBody.Username,
    })
}

func addSessionForUserId(c echo.Context, userId string) error {
    sess, err := session.Get("session", c)
    if err != nil {
        return err
    }

    sess.Options = &sessions.Options{
        Path: "/",
        MaxAge: 86400 * 7,
        HttpOnly: true,
        // set secure to true on prod so cookies are only set for https requests
        // Secure: true,
    }

    // TODO change this to user.ID from the creds check above
    sess.Values["userID"] = userId

    return sess.Save(c.Request(), c.Response())
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

    //TODO check username/password against the db and get the user
    err := addSessionForUserId(c, fmt.Sprintf("%d", getUserID()))
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Couldn't add session in login: %s", err))
    }

    return c.JSON(http.StatusOK, LoginResp{
        Username: loginBody.Username,
    })
}

func LogOut(c echo.Context) error {
    sess, err := session.Get("session", c)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound)
    }

    // MaxAge < 0 == delete imediately
    sess.Options = &sessions.Options{
        Path: "/",
        MaxAge: -1,
        HttpOnly: true,
    }

    if err := sess.Save(c.Request(), c.Response()); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Unable to delete session: %s", err))
	}

    return c.String(http.StatusOK, "Logged out")
}

// getUserID is used to get the next userID from the db, probs just get the one
// created from the user model though later
func getUserID() int {
    // TODO use the db to get the id for the user
    return 1
}
