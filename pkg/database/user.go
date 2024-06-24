package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserModel struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	LastName       string             `bson:"last_name"`
	Email          string             `bson:"email"`
	Password       string             `bson:"password"`
	PurchasedBonds []bondApi.Bond     `bson:"purchased_bonds"`
}

func toUserModel(user *bondApi.User) (UserModel, error) {

	var userId primitive.ObjectID
	var err error

	if user.ID == "" {
		userId = primitive.NewObjectID()
	} else {
		userId, err = primitive.ObjectIDFromHex(user.ID)
	}

	if err != nil {
		return UserModel{}, err
	}

	return UserModel{
		ID:             userId,
		Name:           user.Name,
		LastName:       user.LastName,
		Email:          user.Email,
		Password:       user.Password,
		PurchasedBonds: user.PurchasedBonds,
	}, nil
}

func toUserApiModel(user UserModel) *bondApi.User {
	userId := user.ID.Hex()

	return &bondApi.User{
		ID:             userId,
		Name:           user.Name,
		LastName:       user.LastName,
		Email:          user.Email,
		Password:       user.Password,
		PurchasedBonds: user.PurchasedBonds,
	}

}

func (c *MongoController) CreateUser(u *bondApi.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Adjust timeout as needed
	defer cancel()

	col := c.Db.Collection("users")

	//check if bond already exists

	log.Println("accessing mongo controller")
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now

	newUser, err := toUserModel(u)
	if err != nil {
		log.Println("error while mapping new bond")
		return nil
	}

	result, err := col.InsertOne(ctx, newUser)
	if err != nil {
		//log.Println("error inserting bond: ", err)
		return err
	}

	log.Println("successfully inserted new bond: ", result.InsertedID)
	return nil
}

func (c *MongoController) GetUser(username string) (*bondApi.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Adjust timeout as needed
	defer cancel()

	col := c.Db.Collection("users")

	var user UserModel
	err := col.FindOne(ctx, bson.M{"name": bson.M{"$eq": username}}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("element with username: %s is not found", username)
	} else if err != nil {
		return nil, fmt.Errorf("error fetching element: %w", err)
	}

	result := toUserApiModel(user)
	return result, nil
}
