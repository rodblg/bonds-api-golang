package auth

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

var jwtKey = []byte("my_secret")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// type UserInterface interface {
// 	GetBond(w http.ResponseWriter, r *http.Request)
// 	GetAllBonds(w http.ResponseWriter, r *http.Request)
// 	PublishNewBond(w http.ResponseWriter, r *http.Request)
// 	CreateUser(w http.ResponseWriter, r *http.Request)
// 	GetUser(username string) (*bondApi.User, error)
// }

type AuthController struct {
	User *usecases.UsecasesController
}

func NewAuthController(u *usecases.UsecasesController) *AuthController {
	return &AuthController{
		User: u,
	}
}

func (c *AuthController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", c.Login)
	//r.Get("/bond/buy/{id}", u.BuyBond)
	// r.Get()
	// r.Post("/")
	// r.Post()
	return r
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	username, _, ok := r.BasicAuth()
	if !ok {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := c.User.GetUser(username)
	log.Print(user, err)
}
