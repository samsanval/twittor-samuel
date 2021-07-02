package routes

import (
	"encoding/json"
	"net/http"

	"github.com/samsanval/twittor-samuel/bd"
	"github.com/samsanval/twittor-samuel/models"
)

func ModifyProfile(writer http.ResponseWriter, reader *http.Request) {
	var user models.User
	err := json.NewDecoder(reader.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Wrong Data "+err.Error(), http.StatusBadRequest)
		return
	}
	var status bool
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil {
		http.Error(writer, "Error when trying to modify the register "+err.Error(), http.StatusInternalServerError)
		return
	}
	if status == false {
		http.Error(writer, "The register has not been modified", http.StatusNoContent)
	}
	writer.WriteHeader(http.StatusAccepted)
}
