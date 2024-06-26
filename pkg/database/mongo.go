package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoController struct {
	Db       *mongo.Database
	CollName string
	//Storage bondApi.Storage
}

func NewMongoController(db *mongo.Database, coll string) *MongoController {
	return &MongoController{
		Db:       db,
		CollName: coll,
	}
}

func MongoConnection(database string) (*mongo.Database, error) {
	log.Println("connection to mongo...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@mongodb:27017/"))
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Println("error while connecting...")
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(database)
	log.Println("Successfully connected to MongoDB")

	return db, nil

}
