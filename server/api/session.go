package api

import (
    "fmt"
    "net/http"

    "github.com/flippinberger/djmocracy/repo"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"
)

func Register(c echo.Context) error {
    var loginBody LoginBody
    if err := c.Bind(&loginBody); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to parse params: %s", err))
    }

    djDB, err := getDB()
    if err != nil {
        return err
    }

    userID, err := repo.CreateUser(djDB, loginBody.Username, loginBody.Password)
    if err != nil {
        return err
    }

    err = addSessionForUserId(c, userID)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Couldn't add session in register: %s", err))
    }

    return c.JSON(http.StatusOK, LoginResp{
        Username: loginBody.Username,
        UserID: userID,
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
        // TODO do something with SameSite to fix warnings in console
        // SameSite: ,
        // set secure to true on prod so cookies are only set for https requests
        // Secure: true,
    }

    sess.Values["userID"] = userId

    return sess.Save(c.Request(), c.Response())
}

type LoginBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResp struct {
    UserID string `json:"userId"`
    Username string `json:"username"`
}

func LogIn(c echo.Context) error {
    var loginBody LoginBody
    if err := c.Bind(&loginBody); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to parse params: %s", err))
    }

    djDB, err := getDB()
    if err != nil {
        return err
    }
    user, err := repo.CheckLogin(djDB, loginBody.Username, loginBody.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "username or password incorrect")
    }

    err = addSessionForUserId(c, user.ID.String())
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Couldn't add session in login: %s", err))
    }

    return c.JSON(http.StatusOK, LoginResp{
        UserID: user.ID.Hex(),
        Username: user.Username,
    })
}

// TODO do I need to remove the session from the userID <-> sessionID doc in db?
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
