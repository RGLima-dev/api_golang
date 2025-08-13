package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Defines user repo
type Users struct {
	db *sql.DB
}

// Create a user repo
func NewUserRepo(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) ValidityUserEmail(email string) (models.User, error) {

	row, erro := repository.db.Query("SELECT id,passwd from users WHERE email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	var user models.User

	if row.Next() {
		if erro = row.Scan(&user.Id, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
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

func (repository Users) UpdateUser(userId uint64, userToBeUpdated models.User) (models.User, error) {

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

func (repository Users) DeleteUser(userId uint64) error {

	statement, erro := repository.db.Prepare(
		"DELETE FROM users WHERE id = ?;",
	)
	if erro != nil {
		return erro
	}

	defer statement.Close()

	result, erro := statement.Exec(userId)
	if erro != nil {
		return erro
	}

	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro
	}

	if rowsAffected == 0 {
		return fmt.Errorf("not user with the id:", uint64(userId))
	}

	return nil
}

func (repository Users) FollowUser(userId uint64, followerId uint64) error {
	statement, erro := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userId, followerId); erro != nil {
		return erro
	}

	return nil
}

func (repository Users) UnfollowUser(userId uint64, followerId uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	result, erro := statement.Exec(userId, followerId)
	if erro != nil {
		return erro
	}

	rowsAffected, erro := result.RowsAffected()
	if erro != nil {
		return erro
	}

	if rowsAffected == 0 {
		return fmt.Errorf("you do not follow this user")
	}

	return nil
}

func (repository Users) GetAllFollowersOfUser(followers_of_user uint64) ([]models.User, error) {
	rows, erro := repository.db.Query("SELECT u.nickname, u.email FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?", followers_of_user)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if erro := rows.Scan(&user.Nickname, &user.Email); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	if erro := rows.Err(); erro != nil {
		return nil, erro
	}
	return users, nil

}
