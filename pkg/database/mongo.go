package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection(database, collection string) (*mongo.Database, error) {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@mongo:27017/"))
	if err != nil {
		log.Println("could not connect to the database:", err)
		return nil, err
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database(database)
	return db, nil

}
