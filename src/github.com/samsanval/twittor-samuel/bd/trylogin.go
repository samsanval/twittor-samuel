package bd

import (
	"github.com/samsanval/twittor-samuel/models"
	"golang.org/x/crypto/bcrypt"
)

/* Login */
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := CheckIfExistsUser(email)
	if found == false {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
