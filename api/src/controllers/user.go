package controllers

import "net/http"

// CreateUser insert user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User!"))
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
