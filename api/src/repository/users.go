package repository

import (
	"api/src/models"
	"database/sql"
)

// Defines user repo
type Users struct {
	db *sql.DB
}

// Create a user repo
func NewUserRepo(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO User(Name,Nickname,Email,Password) VALUES(?,?,?,?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()
	result, erro := statement.Exec(user.Name, user.Nickname, user.Email, user.Password)
	if erro != nil {
		return 0, nil
	}
	LastInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(LastInsertedId), nil
}
