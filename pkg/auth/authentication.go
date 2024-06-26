package auth

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rodblg/bonds-api-golang/pkg/bondApi"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
	"golang.org/x/crypto/bcrypt"
)

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

	return r
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {

	email, password, ok := r.BasicAuth()
	if !ok {
		err := errors.New("error extracting basic authentication credentials")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBasicAuth))
		return
	}

	user, err := c.User.GetUser(email)
	if err != nil {
		log.Println("user not found")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrNotFound))
		return
	}

	err = VerifyPassword(password, user.Password)
	if err != nil {
		log.Println("invalid credentials")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBadRequest))
		return
	}

	token, _, err := TokenGenerator(user.Email, user.ID)
	if err != nil {
		log.Println("failed token generation")
		render.Render(w, r, bondApi.ErrRender(err, http.StatusBadRequest, bondApi.ErrBadRequest))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	w.WriteHeader(http.StatusCreated)
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, givenPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	if err != nil {
		log.Println(err)
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
