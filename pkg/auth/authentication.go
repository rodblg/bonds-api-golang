package auth

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
	"golang.org/x/crypto/bcrypt"
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

	email, password, ok := r.BasicAuth()
	if !ok {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := c.User.GetUser(email)
	if err != nil {
		log.Println("user not found", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// 	log.Print(user, err)

	err = VerifyPassword(password, user.Password)
	if err != nil {
		log.Println("invalid credentials")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, refershToken, _ := TokenGenerator(user.Email, user.ID)
	log.Println(refershToken)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	w.WriteHeader(http.StatusCreated)
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		//handle error
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, givenPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	if err != nil {
		//handle error
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
