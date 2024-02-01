package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test/pkg/res"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	body, _ := res.GetBody(r)
	name := body["name"].(string)
	email := body["email"].(string)
	password := body["password"].(string)

	err := InsertUser(name, email, password)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "User created"})
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := SelectAllUsers()
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, users)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	body, _ := res.GetBody(r)
	name := body["name"].(string)
	email := body["email"].(string)
	password := body["password"].(string)

	err = UpdateUser(id, name, email, password)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "User updated"})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	err = DeleteUser(id)
	if err != nil {
		res.JSON(w, r, http.StatusInternalServerError, res.Json{"message": err.Error()})
		return
	}

	res.JSON(w, r, http.StatusOK, res.Json{"message": "User deleted"})
}
