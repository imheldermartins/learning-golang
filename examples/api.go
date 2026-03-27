package examples

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloResponse struct {
	Route   string `json:"route"`
	Message string `json:"message"`
	Method  string `json:"method,omitempty"`
}

const PORT int = 8080
const BASE_URL string = "http://127.0.0.1"

var URL string = fmt.Sprintf("%s:%d", BASE_URL, PORT)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HelloResponse{
		Route:   "/hello",
		Method:  r.Method,
		Message: "Hey user, hello!",
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func Run() {
	log.Printf("Server is running on %s\n", URL)

	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", PORT), nil)
	if err != nil {
		log.Fatal("failed to start server")
	}
}

func Hello() {
	http.HandleFunc("/hello", helloHandler)

	Run()
}
