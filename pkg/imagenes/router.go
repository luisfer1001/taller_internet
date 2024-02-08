package imagenes

import "github.com/go-chi/chi/v5"

//prueba

func NewRouter(r *chi.Mux) {
	r.Post("/imagenes", postImagen)
	r.Get("/imagenes/all", getAllImagenes)
	r.Put("/imagenes/{id}", putImagen)
	r.Delete("/imagenes/{id}", deleteImagen)
}
