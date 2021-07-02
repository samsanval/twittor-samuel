package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/samsanval/twittor-samuel/bd"
	"github.com/samsanval/twittor-samuel/models"
)

func NewTweet(writer http.ResponseWriter, reader *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(reader.Body).Decode(&message)

	register := models.NewTweet{
		UserId:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(writer, "An error has ocurred trying to insert the Tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(writer, "There's been an error inserting the tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
