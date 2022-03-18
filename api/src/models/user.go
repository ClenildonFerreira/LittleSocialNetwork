package models

import (
	"errors"
	"strings"
	"time"
)

// User represents user used in social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name_user,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password_user,omitempty"`
	CreatedIn time.Time `json:"createdIn,omitempty"`
}

// Prepare will call the methods to validate and format the received user
func (user *User) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The name field cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("The nick field cannot be empty")
	}

	if user.Email == "" {
		return errors.New("The email field cannot be empty")
	}

	if user.Password == "" {
		return errors.New("The password field cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
