package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

type UserController struct {
	Usecase *usecases.UsecasesController
}

func NewUserController(usecase *usecases.UsecasesController) *UserController {
	return &UserController{
		Usecase: usecase,
	}
}

func (u *UserController) Routes() chi.Router {
	r := chi.NewRouter()

	//r.Get("/bond/{id}", u.GetBond)
	//r.Get("/bond", u.GetAllBonds)
	r.Post("/bond", u.NewBonds)
	// r.Get()
	// r.Post("/")
	// r.Post()
	return r
}

// func (u *UserController) GetBond() (w http.ResponseWriter, r *http.Request) {

// }

// func (u *UserController) GetAllBonds() (w http.ResponseWriter, r *http.Request) {

// }

func (u *UserController) NewBonds(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error readint request body")
		return
	}
	defer r.Body.Close()

	//decoder := json.NewDecoder(r.Body)
	decoder := json.NewDecoder(bytes.NewReader(body))

	var newData bondApi.Bond

	err = decoder.Decode(&newData)
	if err != nil {
		log.Println("error while unmarshaling into newBond variable: ", err)
		return
	}

	log.Println("Data Received")
	log.Printf("%+v", newData)

	err = u.Usecase.SellNewBonds(newData)
	if err != nil {
		log.Println("error while creating new bonds: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newData)

}
