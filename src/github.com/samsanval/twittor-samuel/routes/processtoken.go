package routes

import (
	"errors"
	"strings"

	jwtGo "github.com/dgrijalva/jwt-go"
	"github.com/samsanval/twittor-samuel/bd"
	"github.com/samsanval/twittor-samuel/models"
)

var Email string
var IDUser string

/* Process Token*/
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("MasterGo")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid Token Format")
	}
	token = strings.TrimSpace(splitToken[1])
	tokenAnswer, err := jwtGo.ParseWithClaims(token, claims, func(token *jwtGo.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		_, found, _ := bd.CheckIfExistsUser(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()

		}
	}
	if !tokenAnswer.Valid {
		return claims, false, string(""), errors.New("Invalid token")
	}
	return claims, false, string(""), err
}
