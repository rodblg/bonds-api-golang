package http

import "github.com/go-chi/chi/v5"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) Routes() chi.Router {
	r := chi.NewRouter()
	// r.Get()
	// r.Post("/")
	// r.Post()
	return r
}
