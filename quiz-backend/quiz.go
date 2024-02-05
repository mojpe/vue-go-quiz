package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// structure of quiz
type Quizes struct {
	Question  	  string   `json:"question,omitempty"`
	Answers   	  *Answers `json:"answers,omitempty"`
	CorrectAnswer string   `json:"correctAnswer,omitempty"`
}
type Answers struct {
	A  string `json:"a,omitempty"`
	B string  `json:"b,omitempty"`
	C string  `json:"c,omitempty"`
	D string  `json:"d,omitempty"`
}

// structure of quiz
type Players struct {
	Name  	  string   `json:"name,omitempty"`
	Score     string   `json:"score,omitempty"`
}

var quiz []Quizes
var players []Players

// Get the Questions and answers
func GetQuiz(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(quiz)
}

// Get the Players and scores
func GetPlayers(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	json.NewEncoder(w).Encode(players)
}

func handleRequests() {

	router := mux.NewRouter()

	// Hardcoded data for quiz
	quiz = append(quiz, Quizes{Question: "Who is the strongest?", Answers: &Answers{A: "Superman", B: "The Terminator", C: "Waluigi, obviously"}, CorrectAnswer: "c"})
	quiz = append(quiz, Quizes{Question: "What is the best site ever created?", Answers: &Answers{A: "SitePoint", B: "Simple Steps Code", C: "Trick question; they're both the best"}, CorrectAnswer: "c"})
	quiz = append(quiz, Quizes{Question: "Where is Waldo really?", Answers: &Answers{A: "Antarctica", B: "Exploring the Pacific Ocean", C: "Sitting in a tree", D: "Minding his own business, so stop asking"}, CorrectAnswer: "d"})
	
	// Hardcoded data for players
	players = append(players, Players{Name: "John Doe", Score: "1"})
	players = append(players, Players{Name: "Tom Hoe", Score: "2"})
	
	router.HandleFunc("/quiz", GetQuiz).Methods("GET")
	router.HandleFunc("/players", GetPlayers).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8888", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Main Function
func main() {
    handleRequests()
}
