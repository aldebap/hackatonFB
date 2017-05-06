package app

import (
	"log"
	"service"
	"strconv"
	"strings"
	"time"
)

func Process(payload string) {
	request := parse(payload)

	if request != nil {
		service.Service(*request)
	}
}

func parse(payload string) *service.Request {
	split := strings.Split(payload, ";")

	if len(split) < 19 {
		return nil
	}

	request := service.Request{
		Id:                   atoi(split[0]),
		Type:                 split[1],
		Date:                 toDate(split[2], "02/01/2006 15:04:05"),
		RejectReason:         split[3],
		SettlementAdjustment: toDate(split[4], "02/01/2006"),
		GrossValue:           atoi(split[5]),
		UserId:               split[6],
		TechnologyType:       atoi(split[7]),
		StatusRequest:        split[8],
		Customer:             split[9],
		ModCustomer:          atoi(split[10]),
		BatchId:              atoi(split[11]),
		MovementType:         atoi(split[12]),
		RefundStatus:         atoi(split[13]),
		AdjustmentReason:     atoi(split[14]),
		AdjustmentComments:   split[15],
		FileId:               atoi(split[16]),
		Product:              atoi(split[17]),
		EntryType:            split[18],
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
