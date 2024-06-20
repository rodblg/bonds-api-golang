package main

import (
	"log"

	"github.com/rodblg/bonds-api-golang/pkg/database"
	http "github.com/rodblg/bonds-api-golang/pkg/http"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

func main() {

	databaseName := ""
	collectionName := ""

	mongoDatabase, err := database.MongoConnection(databaseName, collectionName)
	if err != nil {
		log.Println("error with database connection")
	}

	usecasesController := usecases.NewUsecasesController(*mongoDatabase)

	http.ListenAndServe(usecasesController)
}
