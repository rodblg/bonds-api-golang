package main

import (
	"log"

	"github.com/rodblg/bonds-api-golang/pkg/database"
	http "github.com/rodblg/bonds-api-golang/pkg/http"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

func main() {

	databaseName := "cicada"
	collectionName := "testing"

	mongoDatabase, err := database.MongoConnection(databaseName)
	if err != nil {
		log.Println("error with database connection", err)
	}

	storage := database.NewMongoController(mongoDatabase, collectionName)

	usecasesController := usecases.NewUsecasesController(storage)

	http.ListenAndServe(usecasesController)
}
