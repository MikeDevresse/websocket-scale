package router

import (
	"encoding/json"
	"log"
)

func Handle(payload string) {
	var jsonPayload []string
	err := json.Unmarshal([]byte(payload), &jsonPayload)
	if err != nil {
		log.Printf("Can not parse %s to json", payload)
	}

	log.Printf("%v", jsonPayload)
}
