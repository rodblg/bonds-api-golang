package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/rodblg/bonds-api-golang/pkg/auth"
	"github.com/rodblg/bonds-api-golang/pkg/usecases"
)

func ListenAndServe(usecases *usecases.UsecasesController) {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	//r.Post("/login")
	//r.Mount("/auth", Login.Router())
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write(([]byte("hello world")))
	// })

	UserController := NewUserController(usecases)
	//authController := auth.NewAuthController(UserController)
	authController := auth.NewAuthController(usecases)

	r.Mount("/auth", authController.Routes())
	r.Mount("/user", UserController.Routes())

	http.ListenAndServe(":8080", r)
}
