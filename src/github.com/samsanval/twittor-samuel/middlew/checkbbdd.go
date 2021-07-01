package middlew

import (
	"net/http"

	"github.com/samsanval/twittor-samuel/bd"
)

/* Middleware check database */
func CheckBBDD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Lost Connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
