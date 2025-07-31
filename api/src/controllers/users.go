package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/resps"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resps.ERROR(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		resps.ERROR(w, http.StatusBadRequest, erro)
	}

	if erro = user.PrepareData(); erro != nil {
		resps.ERROR(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	user.Id, erro = repo.Create(user)
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}

	resps.JSON(w, http.StatusCreated, user)

}
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all Users"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting a User"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updtating User"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User"))
}
