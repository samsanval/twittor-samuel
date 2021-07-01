package main

import (
	"log"

	"github.com/samsanval/twittor-samuel/bd"
	"github.com/samsanval/twittor-samuel/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection to DB")
		return
	}
	handlers.Handlers()
}
