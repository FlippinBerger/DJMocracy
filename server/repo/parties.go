package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/flippinberger/djmocracy/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PartiesCollection = "parties"

func CreateParty(djDB *db.DJDB, name string, time time.Time, owner string) (string, error) {
    ownerID, err := primitive.ObjectIDFromHex(owner)
    if err != nil {
        return "", err
    }

    party := &Party {
        Name: name,
        Time: time,
        Songs: make([]Song, 0, 5),
        Creator: ownerID,
        Owners: make([]primitive.ObjectID, 0, 5),
    }

    res, err := djDB.DB.Collection(PartiesCollection).InsertOne(context.Background(), party)
    if err != nil {
        return "", err
    }

    if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
        return oid.Hex(), nil
    }

    return "", fmt.Errorf("couldn't parse ObjectID to hex string")
}
