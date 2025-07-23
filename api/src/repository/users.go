package repository

import (
	"api/src/models"
	"database/sql"
)

// Defines user repo
type users struct {
	db *sql.DB
}

// Create a user repo
func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}
