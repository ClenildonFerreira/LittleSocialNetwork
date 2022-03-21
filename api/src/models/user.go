package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("The name field cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("The nick field cannot be empty")
	}

	if user.Email == "" {
		return errors.New("The email field cannot be empty")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("invalid email")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("The password field cannot be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordWithHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(passwordWithHash)
	}

	return nil
}
