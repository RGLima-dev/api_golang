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
		"INSERT INTO Users(username,nickname,email,passwd) VALUES(?,?,?,?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()
	result, erro := statement.Exec(user.Username, user.Nickname, user.Email, user.Password)
	if erro != nil {
		return 0, nil
	}
	LastInsertedId, erro := result.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(LastInsertedId), nil
}

func (repository Users) GetAllUsers() ([]models.User, error) {
	rows, erro := repository.db.Query(
		"SELECT id, nickname FROM users",
	)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User
		if erro := rows.Scan(&user.Id, &user.Nickname); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	if erro := rows.Err(); erro != nil {
		return nil, erro
	}

	return users, nil

}

func (repository Users) GetSpecifiedUser(userId int) (models.User, error) {
	var user models.User

	erro := repository.db.QueryRow(
		"SELECT id, nickname FROM users WHERE id = ?", userId,
	).Scan(&user.Id, &user.Nickname)

	if erro != nil {
		return models.User{}, erro
	}
	return user, nil
}

func (repository Users) UpdateUser(userId int, userToBeUpdated models.User) (models.User, error) {

	statment, erro := repository.db.Prepare(
		"UPDATE users SET nickname = ? WHERE id = ?")

	if erro != nil {
		return models.User{}, nil
	}
	defer statment.Close()

	_, erro = statment.Exec(userToBeUpdated.Nickname, userId)
	if erro != nil {
		return models.User{}, nil
	}
	userToBeUpdated.Id = uint64(userId)
	return userToBeUpdated, nil
}
