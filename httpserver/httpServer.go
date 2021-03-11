package main

import (
	"QuizMaster/quizjson"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var idFileMap map[string]string = make(map[string]string)

func init() {
	banksDir := `..` + string(os.PathSeparator) + `banks`
	//open parent directory
	files, err := ioutil.ReadDir(banksDir)
	if err != nil {
		log.Fatal(`Error: couldn't read JSON directory`)
	}

	//look for JSON question banks
	for _, fi := range files {
		if strings.Contains(fi.Name(), `.json`) {
			//associate each bank id w/ filename in global map
			filename := banksDir + string(os.PathSeparator) + fi.Name()
			fmt.Println(`Found:`, filename)
			e := quizjson.FromFile(filename)
			idFileMap[e.ID] = filename
		}
	}
	fmt.Println("READY")
}

func info(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	//collect all relevant JSON question banks
	b := quizjson.Banks{}
	for id, file := range idFileMap {
		e := quizjson.FromFile(file)
		b.Banks = append(b.Banks, quizjson.Bank{Name: e.Name, ID: id, Count: len(e.Entries)})
	}

	//return in JSON format
	bytes := b.ToJSON()
	io.WriteString(res, string(bytes))
}

func hello(res http.ResponseWriter, req *http.Request) {
	//read POST request
	defer req.Body.Close()
	msg, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(`Error: couldn't read POST request`)
	}
	jsonReq := quizjson.ToReqJSON(msg)

	//retrieve entries according to request
	q := quizjson.Questions{}
	for i, id := range jsonReq.IDs {
		//get entries from file
		e := quizjson.FromFile(idFileMap[id])

		//apply a random subset of size jsonReq.Count from e.Entries
		selected := rand.Perm(len(e.Entries))[:jsonReq.Count[i]]
		for _, indx := range selected {
			//apply subset
			e.Entries[indx].Question.ShuffleChoices()
			q.Questions = append(q.Questions, e.Entries[indx].Question)
		}
	}

	//shuffle question set
	q.ShuffleQuestions()
	msg = q.ToJSON()
	io.WriteString(res, string(msg))
}

func main() {
	http.HandleFunc(`/req`, hello)
	http.HandleFunc(`/`, info)
	http.ListenAndServe(`:80`, nil)
}
