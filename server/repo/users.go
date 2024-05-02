package repo

import (
	"context"
	"fmt"

	"github.com/flippinberger/djmocracy/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserCollection = "users"

func getCollection(djDB *db.DJDB, coll string) *mongo.Collection {
    return djDB.DB.Collection(coll)
}

func CreateUser(djDB *db.DJDB, username, password string) (string, error) {
    // TODO hash the password before storing it in the db
    user := &User{
        ID: primitive.NewObjectID(),
        Username: username,
        Password: password,
    }
    _, err := getCollection(djDB, UserCollection).InsertOne(context.Background(), user)
    if err != nil {
        return "", err
    }

    return user.ID.Hex(), nil
}

func CheckLogin(djDB *db.DJDB, username, password string) (*User, error) {
    filter := bson.M{"username": username}

    var user User
    err := getCollection(djDB, UserCollection).FindOne(context.Background(), filter).Decode(&user)
    if err != nil {
        fmt.Println("login err is finding with filter", err)
        return nil, err
    }

    // TODO do some hashing on the password to check
    if user.Password == password {
        return &user, nil
    }

    fmt.Println("passwords didn't match")
    return nil, fmt.Errorf("passwords don't match")
}
