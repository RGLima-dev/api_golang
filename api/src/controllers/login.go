package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/resps"
	"api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

// Auth a user
func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		resps.ERROR(w, http.StatusUnprocessableEntity, erro)
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
	userAuth, erro := repo.ValidityUserEmail(user.Email)
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
		return
	}
	if erro = security.ValidyPasswd(userAuth.Password, user.Password); erro != nil {
		resps.JSON(w, http.StatusUnauthorized, map[string]string{"erro": "wrong password"})
		return
	}
	token, erro := auth.CreateToken(userAuth.Id)
	if erro != nil {
		resps.ERROR(w, http.StatusInternalServerError, erro)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.LoginResponse{Login_token: token})
}
