package models

import "time"

// defines a user
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Nickname  string    `json:"Nickname,omitempty"`
	Email     string    `json:"Email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}
