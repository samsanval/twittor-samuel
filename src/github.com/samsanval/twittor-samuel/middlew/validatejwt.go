package middlew

import (
	"net/http"

	"github.com/samsanval/twittor-samuel/routes"
)

/*Validate JWT*/
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, reader *http.Request) {
		_, _, _, err := routes.ProcessToken(reader.Header.Get("Authorization"))
		if err != nil {
			http.Error(writer, "Error in Token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(writer, reader)
	}
}
