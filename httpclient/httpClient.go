package main

import (
	"QuizMaster/quizjson"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func sendRequest(bankIDs []string, qCount []int) {
	//prep and send POST request to https://quizmaster1.azurewebsites.net/req
	req := quizjson.ReqJSON{IDs: bankIDs, Count: qCount}
	bs := req.ToJSON()
	resp, err := http.Post(`https://quizmaster1.azurewebsites.net/req`, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatal(`Error: couldn't POST to https://quizmaster1.azurewebsites.net/req`)
	}
	defer resp.Body.Close()

	//catch and interpret response
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read POST response`)
	}

	//display question set to user
	q := quizjson.ToQuestionSet(msg)
	for _, question := range q.Questions {
		fmt.Println(question)
	}
}

func getInfo() {
	//send POST request to https://quizmaster1.azurewebsites.net/req
	bs := make([]byte, 10)
	resp, err := http.Post(`https://quizmaster1.azurewebsites.net`, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatal(`Error: couldn't POST to https://quizmaster1.azurewebsites.net/req`)
	}
	defer resp.Body.Close()

	//catch response
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read POST response`)
	}

	//display formatted response to user
	b := quizjson.ToBanks(msg)
	for _, bank := range b.Banks {
		fmt.Println(bank)
		fmt.Println(``)
	}
}

func main() {
	//test concurrency
	for i := 0; i < 10; i++ {
		go getInfo()
		go sendRequest([]string{`1`, `2`}, []int{3, 3})
	}

	//send request to HTTP Server
	//getInfo()

	args := os.Args

	if len(args) > 1 {
		fmt.Println("=======================================")
		fmt.Println("httpClient.go sends a request to a server.")
		fmt.Println("Returns back a bank of questions and answers.")
		fmt.Println("=======================================")
		os.Exit(0)
	}
	sendRequest([]string{`x2856m`, `x2856k`, "x2856j"}, []int{3, 2, 3})
	sendRequest([]string{`1`, `2`}, []int{3, 3})
}
