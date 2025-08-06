package models

import (
	"errors"
	"strings"
	"time"
)

// defines a user
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Username  string    `json:"Name,omitempty"`
	Nickname  string    `json:"Nickname,omitempty"`
	Email     string    `json:"Email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	CreatedAt time.Time `json:"-"`
}

func (user *User) PrepareData() error {
	if erro := user.Validity(); erro != nil {
		return erro
	}
	user.format()
	return nil
}

func (user *User) Validity() error {
	if user.Username == "" {
		return errors.New("username is blank")
	}

	if user.Nickname == "" {
		return errors.New("nickname is blank")
	}

	if user.Email == "" {
		return errors.New("email is blank")
	}

	if user.Password == "" {
		return errors.New("password is blank")
	}
	return nil
}

func (user *User) format() {
	user.Username = strings.TrimSpace(user.Username)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}
