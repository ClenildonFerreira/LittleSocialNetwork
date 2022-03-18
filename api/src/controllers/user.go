package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser insert user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyResquest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyResquest, &user); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

// SearchUsers Search all users saved in the database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search all Users!"))
}

// SearchUser Search user saved in the database
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Search User!"))
}

// UpdateUser update a user information saved in database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}

// DeleteUser delete a user information saved in database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User!"))
}
