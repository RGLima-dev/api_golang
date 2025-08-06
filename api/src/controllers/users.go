package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/resps"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	var users []models.User

	db, erro := database.Connect()
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	users, erro = repo.GetAllUsers()
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}

	resps.JSONpretty(w, http.StatusAccepted, users)

}
func GetSpecificUser(w http.ResponseWriter, r *http.Request) {
	id_var := mux.Vars(r)
	idStr := id_var["userId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resps.ERROR(w, http.StatusBadRequest, errors.New("id inválido"))
		return
	}

	var user models.User

	db, erro := database.Connect()
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	user, erro = repo.GetSpecifiedUser(id)
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, errors.New("id not found"))
	}

	resps.JSON(w, http.StatusOK, user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resps.ERROR(w, http.StatusUnprocessableEntity, erro)
		return
	}
	id_var := mux.Vars(r)
	idStr := id_var["userId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resps.ERROR(w, http.StatusBadRequest, errors.New("id inválido"))
		return
	}
	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		resps.ERROR(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	user, erro = repo.UpdateUser(id, user)
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}

	resps.JSON(w, http.StatusCreated, user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resps.ERROR(w, http.StatusUnprocessableEntity, erro)
		return
	}
	id_var := mux.Vars(r)
	idStr := id_var["userId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resps.ERROR(w, http.StatusBadRequest, errors.New("id inválido"))
		return
	}
	var user models.User

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		resps.ERROR(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	erro = repo.DeleteUser(id)
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
	}
	resps.JSON(w, http.StatusAccepted, map[string]interface{}{
		"message": fmt.Sprintf("User with id %d deleted with success! id:", uint64(id)),
	})
}
