package repo

import (
	"context"
	"time"

	"github.com/flippinberger/djmocracy/db"
)

const PartiesCollection = "parties"

func CreateParty(djDB *db.DJDB, name string, time time.Time, owner string) error {
    party := &Party {
        Name: name,
        Time: time,
        Songs: make([]Song, 0, 5),
        Creator: owner,
        Owners: make([]string, 0, 5),
    }

    _, err := djDB.DB.Collection(PartiesCollection).InsertOne(context.Background(), party)
    if err != nil {
        return err
    }

    return nil
}
