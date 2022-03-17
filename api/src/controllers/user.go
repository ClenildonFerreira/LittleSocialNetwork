package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser insert user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyResquest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(bodyResquest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Conect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewRepositoryUser(db)
	repository.Create(user)
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
