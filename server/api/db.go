package api

import (
    "net/http"
	"github.com/labstack/echo/v4"
    "github.com/flippinberger/djmocracy/db"
)

func getDB() (*db.DJDB, error) {
    djDB := db.GetDB()
    if djDB == nil {
        return nil, echo.NewHTTPError(http.StatusInternalServerError, "Unable to get db")
    }

    return djDB, nil
}
