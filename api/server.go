package api

import (
	"fmt"
	"log"
	"net/http"

	"learning/models"
)

const PORT int = 8080
const BASE_URL string = "127.0.0.1"

var URL string = fmt.Sprintf("%s:%d", BASE_URL, PORT)

func Server() {
	models.Init()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", GetUser)

	log.Printf("Server is running on %s\n", URL)

	err := http.ListenAndServe(URL, mux)
	if err != nil {
		log.Fatal("failed to start server")
	}
}