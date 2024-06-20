package usecases

import "go.mongodb.org/mongo-driver/mongo"

type Usecases struct {
	Db mongo.Database
}

func NewUsecasesController(database mongo.Database) *Usecases {
	return &Usecases{
		Db: database,
	}
}
