package user

import "github.com/go-chi/chi/v5"

func NewRouter(r *chi.Mux) {
	r.Post("/user", postUser)
	r.Get("/user/all", getAllUsers)
	r.Put("/user/{id}", putUser)
	r.Delete("/user/{id}", deleteUser)
}
