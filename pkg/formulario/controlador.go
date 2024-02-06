package formulario

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func postFormulario(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	name := body["name"].(string)
	descripcion := body["descripcion"].(string)
	

	err := InsertFormulario(name, descripcion)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "User created"})
}

func getAllFormulario(w http.ResponseWriter, r *http.Request) {
	users, err := SelectAllFormularios()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, users)
}

func putFormulario(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	body, _ := res.GetBody(r)
	name := body["name"].(string)
	descripcion := body["descripcion"].(string)
	

	err = Updateformulario(id, name, descripcion)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "User updated"})
}

func deleteFormulario(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	err = DeleteFormulario(id)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "User deleted"})
}
