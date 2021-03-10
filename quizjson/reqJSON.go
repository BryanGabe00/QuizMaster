package quizjson

import (
	"encoding/json"
	"log"
)

//ReqJSON struct represents one JSON-encoded request
type ReqJSON struct {
	IDs   []string `json:"ids"`
	Count []int    `json:"count"`
}

//ToJSON method returns ReqJSON struct as JSON []byte
func (r ReqJSON) ToJSON() []byte {
	bytes, err := json.Marshal(r)
	if err != nil {
		log.Fatal(`Error: couldn't convert ReqJSON struct to JSON []byte`)
	}
	return bytes
}

//ToReqJSON func returns JSON []byte as ReqJSON struct
func ToReqJSON(bytes []byte) ReqJSON {
	var r ReqJSON
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		log.Fatal(`Error: couldn't convert JSON []byte to ReqJSON struct`)
	}
	return r
}
