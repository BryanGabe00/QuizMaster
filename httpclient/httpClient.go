package main

import (
	"QuizMaster/quizjson"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func sendRequest(bankIDs []string, qCount []int) {
	//prep and send POST request to https://quizmaster1.azurewebsites.net/req
	req := quizjson.ReqJSON{IDs: bankIDs, Count: qCount}
	bs := req.ToJSON()
	resp, err := http.Post(`https://quizmaster1.azurewebsites.net/req`, `application/json`,
		bytes.NewBuffer(bs))
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
	resp, err := http.Post(`https://quizmaster1.azurewebsites.net`, `application/json`,
		bytes.NewBuffer(bs))
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

func testEndpoints() {
	//test concurrency
	for i := 0; i < 10; i++ {
		fmt.Println(`Test #`, i)
		go getInfo()
		go sendRequest([]string{`1`, `2`}, []int{3, 3})
	}
}

func main() {
	//check for `help` argument
	args := os.Args[1:]
	if len(args) > 0 {
		if strings.EqualFold(args[0], `help`) {
			fmt.Println(`=============================================`)
			fmt.Println(`Welcome to Help Page!`)
			fmt.Println(`httpClient.go sends a request to a server.`)
			fmt.Println(`Returns a bank of questions and answers.`)
			fmt.Println("Use \u2018test\u2019 to run concurrency test.")
			fmt.Println(`=============================================`)
		} else if strings.EqualFold(args[0], `test`) {
			fmt.Println(`Performing Concurrency Test`)
			testEndpoints()
			fmt.Println(`Success!`)
		}
		os.Exit(0)
	}

	getInfo()

	//prepare to catch and store user values
	userBanks := make([]string, 0)
	userCts := make([]int, 0)

	var userBank string = `temp`
	var userCt int = 1
	for true {
		//ask for bank ID
		fmt.Print(`Enter a bank ID (END to exit): `)
		fmt.Scanln(&userBank)

		//break if ID = "END"
		if strings.EqualFold(userBank, `END`) {
			break
		}

		//ask for question count
		fmt.Print(`Enter number of questions: `)
		fmt.Scanln(&userCt)

		//save only if the ID isn't "END"
		if !strings.EqualFold(userBank, `END`) {
			userBanks = append(userBanks, userBank)
			userCts = append(userCts, userCt)
		}
	}

	//send request to HTTP Server
	if len(userBanks) > 0 {
		sendRequest(userBanks, userCts)
	}

	/*//testing requests for debugging
	sendRequest([]string{`x2856m`, `x2856k`, "x2856j"}, []int{3, 2, 3})
	sendRequest([]string{`1`, `2`}, []int{3, 3}) */
}
