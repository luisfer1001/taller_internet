package formulario

import "github.com/go-chi/chi/v5"

func NewRouter(r *chi.Mux) {
	r.Post("/formulario", postFormulario)
	r.Get("/formulario/all", getAllFormulario)
	r.Put("/formulario/{id}", putFormulario)
	r.Delete("/formulario/{id}", deleteFormulario)
}