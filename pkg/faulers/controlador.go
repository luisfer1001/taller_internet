package faulers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func postFauler(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	code := body["code"].(int)
	faules := body["faules"].(string)

	err := InsertFauler(code, faules)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Fauler created"})
}

func getAllFaulers(w http.ResponseWriter, r *http.Request) {
	faulers, err := SelectAllFaulers()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, faulers)
}

func putFauler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	body, _ := res.GetBody(r)
	code := body["code"].(int)
	faules := body["faules"].(string)

	err = UpdateFauler(id, code, faules)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Fauler updated"})
}

func deleteFauler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	err = DeleteFauler(id)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Fauler deleted"})
}
