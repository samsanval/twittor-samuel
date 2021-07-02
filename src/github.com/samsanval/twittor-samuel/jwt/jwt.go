package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/samsanval/twittor-samuel/models"
)

/*Generate JWT*/
func GenerateJWT(user models.User) (string, error) {
	key := []byte("MasterGo")
	payload := jwtGo.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"dateBirth": user.DateBirth,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwtGo.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
