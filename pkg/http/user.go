package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

type UserController struct {
	usecase usecases.UsecasesController
}

func NewUserController(usecase usecases.UsecasesController) *UserController {
	return &UserController{}
}

func (u *UserController) Routes() chi.Router {
	r := chi.NewRouter()

	//r.Get("/bond/{id}", u.GetBond)
	//r.Get("/bond", u.GetAllBonds)
	r.Post("/bond", u.NewBond)
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

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var newData bondApi.Bond

	err := decoder.Decode(&newData)
	if err != nil {
		log.Println("error while unmarshaling into newBond variable: ", err)
		return
	}

	log.Println("Data Received")
	log.Printf("%+v", newBond)

	err = usecases.SellNewBonds(newData)
	if err != nil {
		log.Println("error while creating new bonds: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBond)

}
