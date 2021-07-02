package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/samsanval/twittor-samuel/bd"
)

func GetTweets(writer http.ResponseWriter, reader *http.Request) {
	ID := reader.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Missing Id field", http.StatusBadRequest)
		return
	}
	if len(reader.URL.Query().Get("page")) < 1 {
		http.Error(writer, "Missing page field", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(reader.URL.Query().Get("page"))
	if err != nil {
		http.Error(writer, "Page field is wrong "+err.Error(), http.StatusBadRequest)
		return
	}
	pageToInt := int64(page)
	results, correct := bd.GetTweets(ID, pageToInt)
	if correct == false {
		http.Error(writer, "Error finding the results ", http.StatusBadRequest)
		return
	}
	writer.Header().Set("content-type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	json.NewEncoder(writer).Encode(results)
}
