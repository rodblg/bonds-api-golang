package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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
	r.Get("/bond/buy/{id}", u.BuyBond)

	return r
}

func (u *UserController) GetBond(w http.ResponseWriter, r *http.Request) {

	bondId := chi.URLParam(r, "id")
	if bondId == "" {
		log.Println("error extracting id from request")
		err := errors.New("error extracting id from request")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBadRequest))
		return
	}

	//get into usecases
	bond, err := u.Usecase.GetBond(bondId)
	if err != nil {
		log.Println("error fetching bond from db: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}

	//is this to check the availability of the bond??

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bond)
}

func (u *UserController) GetAllBonds(w http.ResponseWriter, r *http.Request) {

	allBonds, err := u.Usecase.GetAllBonds()
	if err != nil {
		log.Println("error fetching all bonds: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allBonds)
}

func (u *UserController) PublishNewBond(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}
	defer r.Body.Close()

	//decoder := json.NewDecoder(r.Body)
	decoder := json.NewDecoder(bytes.NewReader(body))

	var newData bondApi.Bond

	err = decoder.Decode(&newData)
	if err != nil {
		log.Println("error while unmarshaling into newBond variable: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}

	err = u.Usecase.PublishNewBond(newData)
	if err != nil {
		log.Println("error while creating new bonds: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, bondApi.NewBondResponse(&newData))

}

func (u *UserController) BuyBond(w http.ResponseWriter, r *http.Request) {

	//get user info from token auth
	email := r.Context().Value("email").(string)

	user, err := u.Usecase.GetUser(email)
	if err != nil {
		log.Println(err)
		return
	}

	userId := user.ID
	//Get bond request id
	bondId := chi.URLParam(r, "id")
	if bondId == "" {
		log.Printf("missing fetching url %v: ", bondId)
		err := errors.New("missing url from request")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBadRequest))
		return
	}

	bond, err := u.Usecase.GetBond(bondId)
	if err != nil {
		log.Println("error fetching bond: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}

	if bond.Buyer == "" {
		err := u.Usecase.UpdateBondBuyer(bondId, userId)
		if err != nil {
			log.Println("error while updating bond: ", err)
			render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
			return
		}
		err = u.Usecase.UpdateUserBuyedBondsr(userId, *bond)
		if err != nil {
			log.Println("error while updating buyer: ", err)
			render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("the transaction is complete")
	} else {
		// w.Header().Set("Content-Type", "application/json")

		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode("the bond is not available")
		err := errors.New("the bond is not available")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBadRequest))
	}

}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading request body", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBadRequest))
		return
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(bytes.NewReader(body))

	var newData bondApi.User

	err = decoder.Decode(&newData)
	if err != nil {
		log.Println("error while unmarshaling into newUser variable: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}

	newData.Password = auth.HashPassword(newData.Password)

	err = u.Usecase.CreateUser(&newData)
	if err != nil {
		log.Println("error while creating new user: ", err)
		render.Render(w, r, bondApi.ErrRender(err, http.StatusInternalServerError, bondApi.ErrInternalServer))
		return
	}

	// w.WriteHeader(http.StatusCreated)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(newData)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, bondApi.NewUserResponse(&newData))

}

func (u *UserController) GetUser(username string) (*bondApi.User, error) {
	//get into usecases
	user, err := u.Usecase.GetUser(username)
	if err != nil {
		log.Println("error fetching bond: ", err)
		return nil, err
	}

	return user, nil
}
