package quizjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//Entries struct represents a set of JSON-based quiz entries
type Entries struct {
	Name    string  `json:"name"`
	ID      string  `json:"id"`
	Entries []Entry `json:"entries"`
}

//Entry struct represents one JSON-based quiz entry
type Entry struct {
	Question Question `json:"question"`
	Answer   string   `json:"answer"`
}

//String method prints an Entry struct as a string
func (e Entry) String() string {
	retStr := fmt.Sprint(e.Question)
	retStr += "Answer: " + e.Answer
	return fmt.Sprintf(retStr)
}

//FromFile function decodes a .json filepath into an Entries struct
func FromFile(filepath string) Entries {
	//open and read file
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(`Error: couldn't read from JSON file`)
	}

	//convert to Entries struct
	var e Entries
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		log.Fatal(`Error: couldn't convert JSON to Entries struct`)
	}
	return e
}
