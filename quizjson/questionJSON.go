package quizjson

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

//initializer function to seed the random number generator
func init() {
	//seed math/rand for shuffleChoices() and shuffleQuestions() methods
	rand.Seed(time.Now().UnixNano())
}

//Questions struct represents a set of JSON-based quiz questions
type Questions struct {
	Questions []Question `json:"questions"`
}

//Question struct represents one JSON-based quiz question
type Question struct {
	Prompt  string   `json:"prompt"`
	Choices []string `json:"choices"`
}

//String method allows printing Question struct as a string
func (q Question) String() string {
	retStr := q.Prompt + "\n"
	retStr += "============================================\n"
	for _, i := range q.Choices {
		retStr += i + "\n"
	}
	return fmt.Sprintf(retStr)
}

//ShuffleChoices method randomly reorders a slice of q.Choices
func (q *Question) ShuffleChoices() {
	//get a random permuation of indexes
	newIndx := rand.Perm(len(q.Choices))
	newSlice := make([]string, len(q.Choices))

	//create new []string using newIndx
	for i, val := range newIndx {
		newSlice[val] = q.Choices[i]
	}

	q.Choices = newSlice
}

//ToJSON method returns Questions struct as JSON []byte
func (q Questions) ToJSON() []byte {
	bytes, err := json.Marshal(q)
	if err != nil {
		log.Fatal(`Error: couldn't convert Questions struct to JSON`)
	}
	return bytes
}

//ShuffleQuestions method randomly reorders a slice of q.Questions
func (q *Questions) ShuffleQuestions() {
	//get a random permutation of indexes
	newIndx := rand.Perm(len(q.Questions))
	newSlice := make([]Question, len(q.Questions))

	//create new []Question using newIndx
	for i, val := range newIndx {
		newSlice[val] = q.Questions[i]
	}

	q.Questions = newSlice
}

//ToQuestionSet function returns JSON []byte as Questions struct
func ToQuestionSet(bytes []byte) Questions {
	var q Questions
	err := json.Unmarshal(bytes, &q)
	if err != nil {
		log.Fatal(`Error: couldn't convert JSON to Questions struct`, err)
	}
	return q
}
