package repo

import (
	"context"

	"github.com/flippinberger/djmocracy/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoIndices(djdb *db.DJDB) {
    // unique index on username in the Users collection
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"username", 1}},
		Options: options.Index().SetUnique(true),
	}

	djdb.DB.Collection(UserCollection).Indexes().CreateOne(context.Background(), indexModel)
}
