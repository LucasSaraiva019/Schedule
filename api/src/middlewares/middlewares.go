package middlewares

import (
	"log"
	"net/http"
)

func Log(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// func Autenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if erro := authentication.ValidateToken(r); erro != nil {
// 			responses.Erro(w, http.StatusUnauthorized, erro)
// 			return
// 		}
// 		nextFunc(w, r)
// 	}
// }
