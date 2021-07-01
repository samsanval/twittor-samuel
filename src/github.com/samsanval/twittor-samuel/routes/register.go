package routes

import (
	"encoding/json"
	"net/http"

	"github.com/samsanval/twittor-samuel/bd"
	"github.com/samsanval/twittor-samuel/models"
)

/*Register user*/
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request to register "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "The email should be longer than zero", 500)
		return
	}
	if len(user.Password) == 0 {
		http.Error(w, "Password is empty", 500)
		return
	}
	_, found, _ := bd.CheckIfExistsUser(user.Email)
	if found == true {
		http.Error(w, "An user is registered with that email", 400)
		return
	}
	_, status, err := bd.InsertUser(user)
	if err != nil {
		http.Error(w, "An error ocurred", 400)
		return
	}
	if status == false {
		http.Error(w, "An error ocurred", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
