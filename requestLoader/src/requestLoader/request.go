////////////////////////////////////////////////////////////////////////////////
//	request.go  -  May/06/2017  -  aldebaran perseke
//
//	Request file parsing
////////////////////////////////////////////////////////////////////////////////

package requestLoader

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	ID                   int       `json:"id"`
	Type                 string    `json:"type"`
	Date                 time.Time `json:"date"`
	RejectionReason      string    `json:"rejectionReason"`
	SettlementAdjustment time.Time `json:"settlementAdjustment"`
	GrossValue           int       `json:"grossValue"`
	UserID               string    `json:"userId"`
	TechnologyType       int       `json:"technologyType"`
	StatusRequest        string    `json:"statusRequest"`
	Customer             string    `json:"customer"`
	ModCustomer          int       `json:"modCustomer"`
	BatchID              int       `json:"batchId"`
	MovementType         int       `json:"movementType"`
	RefundStatus         int       `json:"refundStatus"`
	AdjustmentReason     int       `json:"adjustmentReason"`
	AdjustmentComments   string    `json:"adjustmentComments"`
	FileID               int       `json:"fileId"`
	Product              int       `json:"product"`
	EntryType            string    `json:"entryType"`
}

var ignoreHeader bool
var verbose bool

////////////////////////////////////////////////////////////////////////////////
//	set the ignore Header flag
////////////////////////////////////////////////////////////////////////////////

func SetIgnoreHeader(_ignoreHeader bool) {

	ignoreHeader = _ignoreHeader
}

////////////////////////////////////////////////////////////////////////////////
//	set the verbose flag
////////////////////////////////////////////////////////////////////////////////

func SetVerbose(_verbose bool) {

	verbose = _verbose
}

////////////////////////////////////////////////////////////////////////////////
//	Parse a CSV into a Request object
////////////////////////////////////////////////////////////////////////////////

func FromCSV(_csvLine string) Request {

	var request Request

	fields := strings.Split(_csvLine, ";")

	if 19 == len(fields) {
		request.ID = atoi(fields[0])
		request.Type = fields[1]
		request.Date = toDate(fields[2], "02/01/2006 15:04:05")
		request.RejectionReason = fields[3]
		request.SettlementAdjustment = toDate(fields[4], "02/01/2006")
		request.GrossValue = atoi(fields[5])
		request.UserID = fields[6]
		request.TechnologyType = atoi(fields[7])
		request.StatusRequest = fields[8]
		request.Customer = fields[9]
		request.ModCustomer = atoi(fields[10])
		request.BatchID = atoi(fields[11])
		request.MovementType = atoi(fields[12])
		request.RefundStatus = atoi(fields[13])
		request.AdjustmentReason = atoi(fields[14])
		request.AdjustmentComments = fields[15]
		request.FileID = atoi(fields[16])
		request.Product = atoi(fields[17])
		request.EntryType = fields[18]
	}

	return request
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

////////////////////////////////////////////////////////////////////////////////
//	Polling the load directory for request files
////////////////////////////////////////////////////////////////////////////////

func LoadRequestFile(_requestFileName string) {

	if true == verbose {
		fmt.Printf("request file found: %s\n", _requestFileName)
	}

	//	open request file
	requestFile, err := os.Open(_requestFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] can't open file %s\n", _requestFileName)
		return
	}

	//	create the line reader
	reader := bufio.NewReader(requestFile)

	//	read all lines from the file
	for lineNumber := 1; ; lineNumber++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			if io.EOF != err {
				fmt.Fprintf(os.Stderr, "[error] reading file %s\n", _requestFileName)
			}

			break
		}

		if 1 < lineNumber || false == ignoreHeader {

			fmt.Printf("parsing record %d\r", lineNumber)

			//	parse the line into a Request object
			var request Request

			request = FromCSV(line)

			SendTopic(request)
		}
	}

	requestFile.Close()
}
