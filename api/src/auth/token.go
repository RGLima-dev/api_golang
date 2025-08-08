package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(uid uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perms["userId"] = uid
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(config.SecretKey)) //secret

}

func ValidateToken(r *http.Request) error {

	tokenStr := extractToken(r)
	token, erro := jwt.Parse(tokenStr, ReturnVerifyKey)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func ReturnVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("signature method incorrect! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenStr := extractToken(r)
	token, erro := jwt.Parse(tokenStr, ReturnVerifyKey)
	if erro != nil {
		return 0, erro
	}

	if perms, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", perms["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return userId, nil
	}
	return 0, errors.New("invalid token")
}
