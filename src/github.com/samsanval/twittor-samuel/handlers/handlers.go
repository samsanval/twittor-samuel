package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/samsanval/twittor-samuel/middlew"
	"github.com/samsanval/twittor-samuel/routes"
)

/*Handlers set port and handler*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckBBDD(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckBBDD(routes.Login)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"

	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
