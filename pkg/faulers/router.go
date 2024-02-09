package faulers

import "github.com/go-chi/chi/v5"

func NewRouter(r *chi.Mux) {
	r.Post("/fauler", postFauler)
	r.Get("/fauler/all", getAllFaulers)
	r.Put("/fauler/{id}", putFauler)
	r.Delete("/fauler/{id}", deleteFauler)
}
