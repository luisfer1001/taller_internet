package imagenes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func postImagen(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	name := body["name"].(string)
	file := body["file"].(string)
	id_formularioStr := body["id_formulario"].(string)

	id_formulario, err := strconv.Atoi(id_formularioStr)
	if err != nil {
		id_formulario = 0
	}

	err = InsertImagen(name, file, id_formulario)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Imagen created"})
}

func getAllImagenes(w http.ResponseWriter, r *http.Request) {
	users, err := SelectAllImagenes()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, users)
}

func putImagen(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	body, _ := res.GetBody(r)
	name := body["name"].(string)
	file := body["file"].(string)
	id_formularioStr := body["id_formulario"].(string)

	id_formulario, err := strconv.Atoi(id_formularioStr)
	if err != nil {
		id_formulario = 0
	}

	err = UpdateImagen(id, name, file, id_formulario)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Imagen updated"})
}

func deleteImagen(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	err = DeleteImagen(id)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "Imagen deleted"})
}
