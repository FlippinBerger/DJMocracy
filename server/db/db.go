package db

import (
	"context"
	"fmt"
    "log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DJDB struct {
    DB *mongo.Database
}

var DB *DJDB

func Connect() error {
    client, err := mongo.Connect(
        context.Background(),
        options.Client().ApplyURI("mongodb://mongo_user:pass@localhost:27017/"),
    )

    if err != nil {
        log.Fatalf("connection error :%v", err)
        return fmt.Errorf("Unable to connect to db: %v", err)
    }

    // TODO do I need to update db name for different environments?
    djDB := client.Database("DJ_db");

    DB = &DJDB{djDB}
    return nil
}

func GetDB() *DJDB {
    return DB
}
