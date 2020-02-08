package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var jwtKey = ""

func getJWTkey() string {
	if jwtKey == "" {
		jwtKey = os.Getenv("JWT_KEY")
		if jwtKey == "" {
			jwtKey = "PengenTinggalDiBandungBrooo"
		}
	}
	return jwtKey
}

func main() {
	router := chi.NewRouter()
	run(router, getJWTkey())

	http.ListenAndServe(":8080", router)
}
