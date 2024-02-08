package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/db"
	"github.com/test/pkg/formulario"
	"github.com/test/pkg/imagenes"
	"github.com/test/pkg/res"
	"github.com/test/pkg/user"
)

func main() {
	err := db.ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return
	}

	createTables()
	httpServer()
}

func httpServer() {
	r := chi.NewRouter()
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		res.JSON(w, r, http.StatusOK, res.Json{"message": "pong"})
	})

	user.NewRouter(r)
	formulario.NewRouter(r)
	imagenes.NewRouter(r)
	log.Println("Server running on port 3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Println(err.Error())
	}
}

func createTables() {
	err := user.CreateTable()
	if err != nil {
		log.Println(err.Error())
	}
	err = formulario.CreateTable()
	if err != nil {
		log.Println(err.Error())
	}
	err = imagenes.CreateTable()
	if err != nil {
		log.Println(err.Error())
	}

}
