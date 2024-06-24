package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rodblg/bonds-api-golang/pkg/auth"
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

	r.Use(auth.Authentication())
	r.Get("/bond/{id}", u.GetBond)
	r.Get("/bond", u.GetAllBonds)
	r.Post("/bond", u.PublishNewBond)

	r.Post("/", u.CreateUser)
	//r.Get("/bond/buy/{id}", u.BuyBond)
	// r.Get()
	// r.Post("/")
	// r.Post()
	return r
}

func (u *UserController) GetBond(w http.ResponseWriter, r *http.Request) {
	//Check user credentials and authorization

	//Get bond request id
	bondId := chi.URLParam(r, "id")
	if bondId == "" {
		//render.Render(w, r, http.StatusBadRequest)
		log.Println("error fetching url {id}")
		return
	}

	log.Println(bondId)
	//get into usecases
	bond, err := u.Usecase.GetBond(bondId)
	if err != nil {
		log.Println("error fetching bond: ", err)
		return
	}

	//is this to check the availability of the bond??

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bond)
}

func (u *UserController) GetAllBonds(w http.ResponseWriter, r *http.Request) {
	//check authorization or that user exists in db
	allBonds, err := u.Usecase.GetAllBonds()
	if err != nil {
		log.Println("error fetching all bonds: ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allBonds)
}

func (u *UserController) PublishNewBond(w http.ResponseWriter, r *http.Request) {
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

	err = u.Usecase.PublishNewBond(newData)
	if err != nil {
		log.Println("error while creating new bonds: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newData)

}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error readint request body")
		return
	}
	defer r.Body.Close()

	//decoder := json.NewDecoder(r.Body)
	decoder := json.NewDecoder(bytes.NewReader(body))

	var newData bondApi.User

	err = decoder.Decode(&newData)
	if err != nil {
		log.Println("error while unmarshaling into newUser variable: ", err)
		return
	}

	newData.Password = auth.HashPassword(newData.Password)
	log.Println(newData)

	err = u.Usecase.CreateUser(&newData)
	if err != nil {
		log.Println("error while creating new user: ", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newData)

}

func (u *UserController) GetUser(username string) (*bondApi.User, error) {
	//Check user credentials and authorization

	//get into usecases
	user, err := u.Usecase.GetUser(username)
	if err != nil {
		log.Println("error fetching bond: ", err)
		return nil, err
	}

	return user, nil
}
