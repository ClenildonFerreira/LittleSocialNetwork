package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Search brings up all users that match a name or nick filter
func (u Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := u.db.Query(
		"select id, name_user, nick, email, createdIn from users where name_user LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// SearchWithID bring a user from the database
func (u Users) SearchWithID(ID uint64) (models.User, error) {
	lines, erro := u.db.Query(
		"select id, name_user, nick, email, createdIn from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); erro != nil {
			return models.User{}, erro
		}

	}

	return user, nil
}
