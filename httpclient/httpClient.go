package main

import (
	"QuizMaster/quizjson"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func sendRequest(bankIDs []string, qCount []int) {
	//prep and send POST request to localhost:80/req
	req := quizjson.ReqJSON{IDs: bankIDs, Count: qCount}
	bs := req.ToJSON()
	resp, err := http.Post(`https://quizmaster1.azurewebsites.net/req`, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatal(`Error: couldn't POST to localhost:80/req`)
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
	//send POST request to localhost:80
	bs := make([]byte, 10)
	resp, err := http.Post(`https://quizmaster1.azurewebsites.net`, `application/json`, bytes.NewBuffer(bs))
	if err != nil {
		log.Fatal(`Error: couldn't POST to localhost:80`)
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
	sendRequest([]string{`1`, `2`}, []int{3, 3})
	//sendRequest([]string{`x2856j`, `x2856k`, `x2856m`}, []int{3, 2, 5})
}
