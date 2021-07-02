package routes

import (
	"encoding/json"
	"net/http"

	"github.com/samsanval/twittor-samuel/bd"
)

func Profile(writer http.ResponseWriter, reader *http.Request) {

	ID := reader.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Id empty", http.StatusBadRequest)
		return
	}
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(writer, "Error trying to look the register "+err.Error(), http.StatusNotFound)
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(writer).Encode(profile)

}
