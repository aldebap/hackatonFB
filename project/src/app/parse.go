package app

import (
	"encoding/json"
	"log"
	"service"
	"strconv"
	"time"
)

func Process(payload string) {
	request := parse(payload)

	if request != nil {
		service.Service(*request)
	}
}

func parse(payload string) *service.Request {
	request := service.Request{}

	err := json.Unmarshal([]byte(payload), &request)
	if err != nil {
		log.Printf("Unsupported payload [%s]: %v", payload, err)

		return nil
	}

	return &request
}

func atoi(data string) int {
	val, err := strconv.Atoi(data)
	if err != nil {
		log.Printf("Error Atoi: %v", data)
	}

	return val
}

func toDate(data string, mask string) time.Time {
	date, err := time.Parse(mask, data)
	if err != nil {
		log.Printf("Error parsing Date %v: %v", data, err)
	}

	return date
}
