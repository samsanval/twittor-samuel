package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samsanval/twittor-samuel/bd"
	"github.com/samsanval/twittor-samuel/jwt"
	"github.com/samsanval/twittor-samuel/models"
)

/* Login */
func Login(writer http.ResponseWriter, reader *http.Request) {
	writer.Header().Add("content-type", "application/json")
	var user models.User
	err := json.NewDecoder(reader.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Invalid User and/or Password"+err.Error(), 500)
		return
	}
	if len(user.Email) == 0 {
		http.Error(writer, "Email is empty", 500)
		return
	}
	document, exists := bd.TryLogin(user.Email, user.Password)
	if exists == false {
		http.Error(writer, "Invalid User and/or Password", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(writer, "Invalid Token Generation "+err.Error(), 400)
		return
	}
	response := models.LoginAnswer{
		Token: jwtKey,
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
