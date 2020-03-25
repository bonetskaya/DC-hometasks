package app

import (
	"fmt"
	//jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	fmt.Println("yay")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
