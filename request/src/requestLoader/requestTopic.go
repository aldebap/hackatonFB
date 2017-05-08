////////////////////////////////////////////////////////////////////////////////
//	requestTopic.go  -  May/07/2017  -  aldebaran perseke
//
//	Request Kafka topic
////////////////////////////////////////////////////////////////////////////////

package requestLoader

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	samara "gopkg.in/Shopify/sarama.v1"
)

var (
	channel  = make(chan string, 50000)
	counter  = int32(0)
	brokers  []string
	config   *samara.Config
	producer samara.AsyncProducer
)

type Process func(string)

func Init(broker []string) {
	brokers = broker

	config = samara.NewConfig()
	config.ClientID = "ProducerRaw"
	config.Consumer.Return.Errors = true
	config.Consumer.Fetch.Min = 1
	config.Producer.Flush.MaxMessages = 100000

	producer, _ = samara.NewAsyncProducer(brokers, config)
}

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

var verbose bool
var topicName string

////////////////////////////////////////////////////////////////////////////////
//	set the verbose flag
////////////////////////////////////////////////////////////////////////////////

func SetVerbose(_verbose bool) {

	verbose = _verbose
}

////////////////////////////////////////////////////////////////////////////////
//	set the Kafka topic name
////////////////////////////////////////////////////////////////////////////////

func SetTopicName(_topicName string) {

	topicName = _topicName
}

////////////////////////////////////////////////////////////////////////////////
//	Polling the load directory for request files
////////////////////////////////////////////////////////////////////////////////

func SendTopic(_request Request) {

	if true == verbose {
		fmt.Printf("[debug] topicName: %s\n", topicName)
	}

	jsonRequest, err := json.Marshal(_request)
	if err != nil {
		log.Panicf("ERROR! %v", err)
	} else {
		producer.Input() <- &samara.ProducerMessage{
			Topic: topicName,
			Value: samara.ByteEncoder(jsonRequest),
			Key:   samara.StringEncoder(strconv.Itoa(_request.ID)),
		}
	}
}
