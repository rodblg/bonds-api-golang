package main

import (
	"log"
	"os"

	"github.com/rodblg/bonds-api-golang/pkg/auth"
	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"github.com/rodblg/bonds-api-golang/pkg/database"
	http "github.com/rodblg/bonds-api-golang/pkg/http"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

func main() {

	USER_INIT_NAME := os.Getenv("USER_INIT_NAME")
	USER_INIT_LASTNAME := os.Getenv("USER_INIT_LASTNAME")
	USER_INIT_PASS := os.Getenv("USER_INIT_PASS")
	USER_INIT_EMAIL := os.Getenv("USER_INIT_EMAIL")

	DATABASE_NAME := os.Getenv("DATABASE_NAME")
	COLLECTION_NAME := os.Getenv("COLLECTION_NAME")

	mongoDatabase, err := database.MongoConnection(DATABASE_NAME)
	if err != nil {
		log.Printf("error with database connection: %v", err)
		log.Fatal(err)
	}

	storage := database.NewMongoController(mongoDatabase, COLLECTION_NAME)

	usecasesController := usecases.NewUsecasesController(storage)
	initialMongoSetUp(usecasesController, USER_INIT_EMAIL, USER_INIT_NAME, USER_INIT_LASTNAME, USER_INIT_PASS)

	http.ListenAndServe(usecasesController)
}

func initialMongoSetUp(u *usecases.UsecasesController, email, name, lastname, pass string) {
	user, err := u.GetUser(email)
	if err != nil {
		log.Printf("error with user setup: %v", err)
	}
	if user == nil {
		password := auth.HashPassword(pass)
		initialUser := bondApi.User{
			Name:     name,
			LastName: lastname,
			Email:    email,
			Password: password,
		}
		err := u.CreateUser(&initialUser)
		if err != nil {
			log.Printf("error with user setup: %v", err)
		}
	}
}
