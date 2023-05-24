package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/araquach/apiAuth/helpers"
	"github.com/araquach/apiAuth/models"
	"github.com/araquach/dbService"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var error helpers.Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		helpers.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		helpers.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		error.Message = "Error generating the password hash."
		helpers.RespondWithError(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = string(hash)

	db.DB.Create(&user)
	if err != nil {
		error.Message = "Server error."
		helpers.RespondWithError(w, http.StatusInternalServerError, error)
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		error.Message = "Error generating token."
		helpers.RespondWithError(w, http.StatusInternalServerError, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	user.Token = token
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	helpers.ResponseJSON(w, user)
}

func GenerateToken(user models.User) (string, error) {
	var err error
	secret := "secret"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "course",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	var error helpers.Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing."
		helpers.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing."
		helpers.RespondWithError(w, http.StatusBadRequest, error)
		return
	}

	password := user.Password

	db.DB.Where("email", user.Email).First(&user)

	hashedPassword := user.Password

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		error.Message = "Invalid Login attempt - please try again"
		helpers.RespondWithError(w, http.StatusUnauthorized, error)
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		error.Message = "Error generating token."
		helpers.RespondWithError(w, http.StatusInternalServerError, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	user.Token = token
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	helpers.ResponseJSON(w, user)
}
