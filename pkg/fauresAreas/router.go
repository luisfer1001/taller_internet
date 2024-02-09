package fauresAreas

import "github.com/go-chi/chi/v5"

func NewRouter(r *chi.Mux) {
	r.Post("/faure", postFaure)
	r.Get("/faure/all", getAllFaures)
	r.Put("/faure/{id}", putFaure)
	r.Delete("/faure/{id}", deleteFaure)
}
