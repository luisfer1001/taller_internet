package fauresAreas

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func postFaure(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	idFauler := body["idFauler"].(string)
	nameArea := body["nameArea"].(string)

	err := InsertFaures(idFauler, nameArea)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Faure created"})
}

func getAllFaures(w http.ResponseWriter, r *http.Request) {
	Faures, err := selectFauresAreas()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, Faures)
}

func putFaure(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	body, _ := res.GetBody(r)
	idFauler := body["idFauler"].(string)
	nameArea := body["nameArea"].(string)

	err = UpdateFaures(id, idFauler, nameArea)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Faure updated"})
}

func deleteFaure(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	err = DeleteFaures(id)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Faure deleted"})
}
