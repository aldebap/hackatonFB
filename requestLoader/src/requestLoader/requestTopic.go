////////////////////////////////////////////////////////////////////////////////
//	requestTopic.go  -  May/06/2017  -  aldebaran perseke
//
//	Request Kafka topic
////////////////////////////////////////////////////////////////////////////////

package requestLoader

import (
	"encoding/json"
	"log"
	"strconv"

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

var topicName string

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
