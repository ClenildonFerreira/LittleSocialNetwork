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

// Create inset a user in database
func (u Users) Create(user models.User) (uint64, error) {
	statement, erro := u.db.Prepare(
		"insert into users (name_user, nick, email, password_user) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	expect, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	LastID, erro := expect.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(LastID), nil
}
