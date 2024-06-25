package main

import (
	"log"

	"github.com/rodblg/bonds-api-golang/pkg/auth"
	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"github.com/rodblg/bonds-api-golang/pkg/database"
	http "github.com/rodblg/bonds-api-golang/pkg/http"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

func main() {

	databaseName := "cicada"
	collectionName := "testing"

	mongoDatabase, err := database.MongoConnection(databaseName)
	if err != nil {
		log.Printf("error with database connection: %v", err)
		log.Fatal(err)
	}

	storage := database.NewMongoController(mongoDatabase, collectionName)

	usecasesController := usecases.NewUsecasesController(storage)
	initialMongoSetUp(usecasesController)

	http.ListenAndServe(usecasesController)
}

func initialMongoSetUp(u *usecases.UsecasesController) {
	user, err := u.GetUser("rb12@email.com")
	if err != nil {
		log.Printf("error with user setup: %v", err)
	}
	if user == nil {
		password := auth.HashPassword("testing")
		initialUser := bondApi.User{
			Name:     "Rodrigo",
			LastName: "Blancas",
			Email:    "rb12@email.com",
			Password: password,
		}
		err := u.CreateUser(&initialUser)
		if err != nil {
			log.Printf("error with user setup: %v", err)
		}
	}
}
