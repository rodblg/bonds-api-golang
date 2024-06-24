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

type BondModel struct {
	ID                       primitive.ObjectID `bson:"_id"`
	Name                     string             `bson:"name"`
	FaceValue                float64            `bson:"face_value"`
	CurrentValue             float64            `bson:"current_value"`
	Isin                     string             `bson:"isin"`
	Issuer                   string             `bson:"issuer"`
	InterestRate             float64            `bson:"interest_rate"`
	InterestPaymentFrequency string             `bson:"interest_payment_frequency"`
	MaturityDate             time.Time          `bson:"maturity_rate"`
	Description              string             `bson:"description"`
	Buyer                    string             `bson:"buyer"`
	CreationAt               time.Time          `bson:"creation_at"`
	UpdatedAt                time.Time          `bson:"updated_at"`
}

func toBondModel(bond bondApi.Bond) (BondModel, error) {

	var bondId primitive.ObjectID
	var err error

	if bond.ID == "" {
		bondId = primitive.NewObjectID()
	} else {
		bondId, err = primitive.ObjectIDFromHex(bond.ID)
	}

	if err != nil {
		return BondModel{}, err
	}

	return BondModel{
		ID:                       bondId,
		Name:                     bond.Name,
		FaceValue:                bond.FaceValue,
		CurrentValue:             bond.CurrentValue,
		Isin:                     bond.Isin,
		Issuer:                   bond.Issuer,
		InterestRate:             bond.InterestRate,
		InterestPaymentFrequency: bond.InterestPaymentFrequency,
		MaturityDate:             bond.MaturityDate,
		Description:              bond.Description,
		Buyer:                    bond.Buyer,
		CreationAt:               bond.CreatedAt,
		UpdatedAt:                bond.UpdatedAt,
	}, nil
}

func toBondApiModel(bond BondModel) *bondApi.Bond {

	bondId := bond.ID.Hex()

	return &bondApi.Bond{
		ID:                       bondId,
		Name:                     bond.Name,
		FaceValue:                bond.FaceValue,
		CurrentValue:             bond.CurrentValue,
		Isin:                     bond.Isin,
		Issuer:                   bond.Issuer,
		InterestRate:             bond.InterestRate,
		InterestPaymentFrequency: bond.InterestPaymentFrequency,
		MaturityDate:             bond.MaturityDate,
		Description:              bond.Description,
		CreatedAt:                bond.CreationAt,
		UpdatedAt:                bond.UpdatedAt,
	}

}

func (c *MongoController) GetBond(bondId string) (*bondApi.Bond, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := c.Db.Collection(c.CollName)

	id, err := primitive.ObjectIDFromHex(bondId)
	if err != nil {
		return nil, fmt.Errorf("error with bondid")
	}

	var bond BondModel
	err = col.FindOne(ctx, bson.M{"_id": bson.M{"$eq": id}}).Decode(&bond)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("element with ID: %s is not found", bondId)
	} else if err != nil {
		return nil, fmt.Errorf("error fetching element: %w", err)
	}

	result := toBondApiModel(bond)

	return result, nil
}

func (c *MongoController) GetAllBonds() ([]bondApi.Bond, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	col := c.Db.Collection(c.CollName)

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error fetching documents")
	}

	defer cursor.Close(ctx)

	var allMongoBonds []BondModel

	if err = cursor.All(context.TODO(), &allMongoBonds); err != nil {
		return nil, fmt.Errorf("error while unpacking cursors into slice")
	}

	var allBonds []bondApi.Bond
	for _, result := range allMongoBonds {
		cursor.Decode(&result)
		newBond := toBondApiModel(result)
		allBonds = append(allBonds, *newBond)
	}
	return allBonds, nil
}

func (c *MongoController) InsertNewBond(b bondApi.Bond) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // Adjust timeout as needed
	defer cancel()

	col := c.Db.Collection(c.CollName)

	//check if bond already exists

	log.Println("accessing mongo controller")
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now

	newBond, err := toBondModel(b)
	if err != nil {
		log.Println("error while mapping new bond")
		return nil
	}

	result, err := col.InsertOne(ctx, newBond)
	if err != nil {
		//log.Println("error inserting bond: ", err)
		return err
	}

	log.Println("successfully inserted new bond: ", result.InsertedID)
	return nil
}

// type Bond struct {
// 	ID                       string    `json:"id,omitempty"`
// 	Name                     string    `json:"name"`
// 	FaceValue                float64   `json:"face_value"`    //Original principal amount of the bond
// 	CurrentValue             float64   `json:"current_value"` //
// 	Isin                     string    `json:"isin"`          //International Securities Identification Number (ISIN)
// 	Issuer                   string    `json:"issuer"`        //Issuer of the bond
// 	InterestRate             float64   `json:"interest_rate"`
// 	InterestPaymentFrequency string    `json:"interest_payment_frequency"`
// 	MaturityDate             time.Time `json:"maturity_date"`
// 	Description              string    `json:"description"`
// }
