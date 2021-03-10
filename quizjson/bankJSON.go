package quizjson

import (
	"encoding/json"
	"fmt"
	"log"
)

//Banks struct represents a set of JSON-based quiz files
type Banks struct {
	Banks []Bank `json:"banks"`
}

//Bank struct represents one JSON-based quiz file
type Bank struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Count int    `json:"Count"`
}

//String method allows printing Bank struct as a string
func (b Bank) String() string {
	retStr := "Bank Name: %s"
	retStr += "\nBank ID: %s"
	retStr += "\nQuestion Count: %d"
	return fmt.Sprintf(retStr, b.Name, b.ID, b.Count)
}

//ToJSON method returns a Banks struct as a JSON []byte
func (b Banks) ToJSON() []byte {
	bytes, err := json.Marshal(b)
	if err != nil {
		log.Fatal(`Error: couldn't convert Banks struct to JSON []byte`)
	}
	return bytes
}

//ToBanks function returns a JSON []byte as a Banks struct
func ToBanks(bytes []byte) Banks {
	var b Banks
	err := json.Unmarshal(bytes, &b)
	if err != nil {
		log.Fatal(`Error: couldn't convert JSON []byte to Banks struct`)
	}
	return b
}
