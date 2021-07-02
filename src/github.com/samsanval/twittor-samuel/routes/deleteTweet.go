package routes

import (
	"net/http"

	"github.com/samsanval/twittor-samuel/bd"
)

func DeleteTweet(writer http.ResponseWriter, reader *http.Request) {
	ID := reader.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Missing ID field", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweet(ID, IDUser)
	if err != nil {
		http.Error(writer, "Something happened when deleting the tweet "+err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
}
