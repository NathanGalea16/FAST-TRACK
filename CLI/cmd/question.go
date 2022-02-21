/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var questionsCmd = &cobra.Command{
	Use:   "question",
	Short: "Fetches a list of questions",
	Long:  `Fetches a question from APi and asks client to input value. the values are then validated from the server and a score is given`,
	Run: func(cmd *cobra.Command, args []string) {
		getQuestions()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)
}

type Answer struct {
	Id  string `json:"id"`
	Ans string `json:"ans"`
}
type Question struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Choice1 string `json:"choice1"`
	Choice2 string `json:"choice2"`
	Choice3 string `json:"choice3"`
	Choice4 string `json:"choice4"`
}

var listOfAllQuestions []Question
var ansArray []Answer

func getQuestions() {
	url := "http://localhost:8080/questions"
	responseBytes := getQuestionData(url)
	question := []Question{}
	if err := json.Unmarshal(responseBytes, &question); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}
	//We got our list of questions now we can store them locally in the cli
	listOfAllQuestions = question
	//We have our questions stored so now we ask the questions
	AskClientAllQuestions(listOfAllQuestions)
}

func AskClientAllQuestions(questionsArray []Question) {
	log.Printf("AskClientAllQuestions")
	fmt.Println(questionsArray)
	ansArray = nil
	for index, ques := range questionsArray {
		//Inits
		var ans Answer
		//Output
		fmt.Printf("Question %v", index)
		fmt.Println(": " + string(ques.Title))
		fmt.Println("	1) " + string(ques.Choice1))
		fmt.Println("	2) " + string(ques.Choice2))
		fmt.Println("	3) " + string(ques.Choice3))
		fmt.Println("	4) " + string(ques.Choice4))

		fmt.Print("\n Your Ans: ")
		//Awaits input from the user
		var choice = ""
		fmt.Scanln(&choice)

		ans.Id = strconv.Itoa(index)
		ans.Ans = choice
		//Stores ans in array
		ansArray = append(ansArray, ans)
	}
	fmt.Printf("Answer Selections - %v ", ansArray)
	ValidateQuestions(ansArray)
}

func ValidateQuestions(ansArray []Answer) {
	log.Println("ValidateQuestions")
	requestBody, err := json.Marshal(ansArray)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://localhost:8080/validate", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response = string(body)
	var responseTrimmed = strings.Replace(response, "[", "", -1)
	responseTrimmed = strings.Replace(responseTrimmed, "]", "", -1)
	responseTrimmed = strings.Replace(responseTrimmed, "\"", "", -1)
	var responseSplit = strings.Split(responseTrimmed, ",")

	fmt.Printf("You have answered %v correctly out of %v questions\n", responseSplit[0], strconv.Itoa(len(listOfAllQuestions)))
	fmt.Printf("You done better than %v%% of users who have done this quiz", responseSplit[1])

}

func getQuestionData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not request a Question")
	}
	request.Header.Add("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response - %v", err)
	}
	return responseBytes
}
