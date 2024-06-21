package usecases

import "go.mongodb.org/mongo-driver/mongo"

type UsecasesController struct {
	Db mongo.Database
}

func NewUsecasesController(database mongo.Database) *UsecasesController {
	return &UsecasesController{
		Db: database,
	}
}
