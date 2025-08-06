package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// defines a user
type User struct {
	Id        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"-"`
}

func (user *User) PrepareData() error {
	if erro := user.Validity(); erro != nil {
		return erro
	}
	if erro := user.format("register"); erro != nil {
		return erro
	}
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
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email invalido")
	}

	if err := checkmail.ValidateHost(user.Email); err != nil {
		return errors.New("e-mail invalido - host inexistente")
	}

	if user.Password == "" {
		return errors.New("password is blank")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashedpasswd, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(hashedpasswd)
	}

	if step == "login" {

	}
	return nil
}
