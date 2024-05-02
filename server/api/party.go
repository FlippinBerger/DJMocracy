package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/flippinberger/djmocracy/repo"
	"github.com/labstack/echo/v4"
)

type Party struct {
    Name string `json:"name"`
    Time time.Time `json:"time"`
    Creator string `json:"creator"`
}

type CreatePartyRes struct {
    PartyID string `json:"partyId"`
}
    
func CreateParty(c echo.Context) error {
    var party Party
    if err := c.Bind(&party); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to parse params: %s", err))
    }

    djDB, err := getDB()
    if err != nil {
        return err
    }

    partyId, err := repo.CreateParty(djDB, party.Name, party.Time, party.Creator)
    if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to create Party: %s", err))
    }

    return c.JSON(http.StatusOK, CreatePartyRes{
        PartyID: partyId,
    })
}
