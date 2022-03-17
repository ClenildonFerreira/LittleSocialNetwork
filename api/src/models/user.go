package models

import "time"

// User represents user used in social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name_user,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password_user,omitempty"`
	CreatedIn time.Time `json:"createdIn,omitempty"`
}
