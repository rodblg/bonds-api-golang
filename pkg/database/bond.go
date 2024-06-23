package database

import (
	"context"
	"log"
	"time"

	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		CreationAt:               bond.CreatedAt,
		UpdatedAt:                bond.UpdatedAt,
	}, nil
}

func (c *MongoController) InsertNewData(b bondApi.Bond) error {
	col := c.Db.Collection(c.CollName)

	log.Println("accessing mongo controller")
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now

	newBond, err := toBondModel(b)
	if err != nil {
		log.Println("error while mapping new bond")
		return nil
	}

	result, err := col.InsertOne(context.Background(), newBond)
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
