// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type QuestionAnswers struct {
	Id  string `json:"id"`
	Ans string `json:"ans"`
}

// Question - Our struct for all Questions and ans
type Question struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Choice1 string `json:"choice1"`
	Choice2 string `json:"choice2"`
	Choice3 string `json:"choice3"`
	Choice4 string `json:"choice4"`
}

type User struct {
	Id    int `json:"Id"`
	Score int `json:"Score"`
}

var Questions []Question
var Server_Answers []QuestionAnswers
var allUsers []User

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	json.NewEncoder(w).Encode(Questions)
}

var ansArray = []QuestionAnswers{} //Temporary array holding answers submitted

func validateQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: validateQuestions")
	ansArray = nil
	w.Header().Set("Content-Type", "application/json")
	var answers []QuestionAnswers
	_ = json.NewDecoder(r.Body).Decode(&answers)
	ansArray = answers

	var score = 0
	for index, ans := range ansArray {
		if ans.Id == Server_Answers[index].Id && ans.Ans == Server_Answers[index].Ans {
			score++
		}
	}
	var percent = getGlobalScores(score)

	var returnValues = []string{strconv.Itoa(score), strconv.Itoa(percent)}
	json.NewEncoder(w).Encode(returnValues)
}

func getGlobalScores(score int) int {
	fmt.Println("Endpoint Hit: getGlobalScores")
	var totalUsers = len(allUsers)
	var greaterScoreCounter = 0
	for _, user := range allUsers {
		if score > user.Score {
			greaterScoreCounter = greaterScoreCounter + 1
		}
	}
	//How many did you outperform

	if totalUsers > 0 {
		percentage := (float64(greaterScoreCounter) / float64(totalUsers)) * 100
		return int(percentage)
	} else {
		return 100
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/questions", returnAllQuestions)
	myRouter.HandleFunc("/validate", validateQuestions).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	Questions = []Question{
		{
			Id:      "0",
			Title:   "Which of these products is sold by the brands Colgate, Oral-B and Sensodyne?",
			Choice1: "Deodorant",
			Choice2: "Shampoo",
			Choice3: "Toothpaste",
			Choice4: "Sun cream",
		}, {
			Id:      "1",
			Title:   "Which tool was used as a weapon by the Norse god Thor?",
			Choice1: "Pliers",
			Choice2: "Hammer",
			Choice3: "Screwdriver",
			Choice4: "Saw",
		},
		{
			Id:      "2",
			Title:   "What is the name of the classic dessert consisting of sponge cake and ice cream covered in meringue?",
			Choice1: "Baked Rhode Island",
			Choice2: "Baked Wyoming",
			Choice3: "Baked Connecticut",
			Choice4: "Baked Alaska",
		},
		{
			Id:      "3",
			Title:   "Trigonometry is a branch of which subject?",
			Choice1: "Biology",
			Choice2: "Economics",
			Choice3: "Psychology",
			Choice4: "Mathematics",
		},
		{
			Id:      "4",
			Title:   "Lily Savage was a persona of which TV personality?",
			Choice1: "Paul O'Grady",
			Choice2: "Barry Humphries",
			Choice3: "Les Dawson",
			Choice4: "Brendan O'Carroll",
		},
		{
			Id:      "5",
			Title:   "Which of these means a speech in a play where a character talks to themselves rather than to other characters?",
			Choice1: "Interlude",
			Choice2: "Revue",
			Choice3: "Soliloquy",
			Choice4: "Chorus",
		},
		{
			Id:      "6",
			Title:   "Which of these is a religious event celebrated in Hinduism?",
			Choice1: "Diwali",
			Choice2: "Ramadan",
			Choice3: "Hanukkah",
			Choice4: "Whitsun",
		},
		{
			Id:      "7",
			Title:   "British athlete Katarina Johnson-Thompson became a world champion in which athletics event in 2019?",
			Choice1: "Heptathlon",
			Choice2: "Marathon",
			Choice3: "100 metres",
			Choice4: "400 metres hurdles",
		},
		{
			Id:      "8",
			Title:   "Which iconic horror film involves a couple whose newborn child is replaced at birth with the Antichrist?",
			Choice1: "The Shining",
			Choice2: "Don't Look Now",
			Choice3: "The Exorcist",
			Choice4: "The Omen",
		},
		{
			Id:      "9",
			Title:   "In the opera by Rossini, what is the first name of The Barber of Seville",
			Choice1: "Tamino",
			Choice2: "Alfredo",
			Choice3: "Don Carlos",
			Choice4: "Figaro",
		},
	}

	Server_Answers = []QuestionAnswers{
		{
			Id:  "0",
			Ans: "3",
		},
		{
			Id:  "1",
			Ans: "2",
		},
		{
			Id:  "2",
			Ans: "4",
		},
		{
			Id:  "3",
			Ans: "4",
		},
		{
			Id:  "4",
			Ans: "1",
		},
		{
			Id:  "5",
			Ans: "3",
		},
		{
			Id:  "6",
			Ans: "1",
		},
		{
			Id:  "7",
			Ans: "1",
		},
		{
			Id:  "8",
			Ans: "4",
		},
		{
			Id:  "9",
			Ans: "4",
		},
	}

	allUsers = []User{
		{
			Id:    0,
			Score: 7,
		},
		{
			Id:    1,
			Score: 8,
		},
		{
			Id:    2,
			Score: 3,
		},
		{
			Id:    3,
			Score: 4,
		},
		{
			Id:    4,
			Score: 5,
		},
		{
			Id:    5,
			Score: 3,
		},
		{
			Id:    6,
			Score: 6,
		},
		{
			Id:    7,
			Score: 6,
		},
		{
			Id:    8,
			Score: 7,
		},
		{
			Id:    9,
			Score: 1,
		},
		{
			Id:    10,
			Score: 10,
		},
		{
			Id:    11,
			Score: 9,
		},
	}
	handleRequests()
}
