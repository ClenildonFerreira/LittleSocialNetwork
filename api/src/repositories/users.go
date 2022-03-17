package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

// NewRepositoryUser Create a users repository
func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{db}
}

func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
