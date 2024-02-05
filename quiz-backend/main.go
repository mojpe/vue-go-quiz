package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// Question structure
type Question struct {
    ID      int      `json:"id"`
    Title   string   `json:"title"`
    Answers []string `json:"answers"`
}

// GetQuestions returns dummy quiz questions
func GetQuestions(w http.ResponseWriter, r *http.Request) {
    questions := []Question{
        {ID: 1, Title: "What is the capital of France?", Answers: []string{"Paris", "Berlin", "London"}},
        {ID: 2, Title: "Which planet is known as the Red Planet?", Answers: []string{"Mars", "Venus", "Earth"}},
        // Add more questions as needed
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(questions)
}

func main() {
	router := mux.NewRouter()

    // Define routes
    router.HandleFunc("/api/questions", GetQuestions).Methods("GET")

    // Enable CORS
    cors := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),  // Adjust based on your needs
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Start the server with CORS middleware
    http.Handle("/", cors(router))
    http.ListenAndServe(":8081", nil)
}
