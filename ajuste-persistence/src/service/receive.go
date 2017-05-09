package service

import (
	"log"
	"time"

	"github.com/wvanbergen/kafka/consumergroup"
	"github.com/wvanbergen/kazoo-go"

	"encoding/json"

	samara "gopkg.in/Shopify/sarama.v1"
)

var (
	channel        = make(chan string, 50000)
	counter        = int32(0)
	brokers        []string
	config         *consumergroup.Config
	producer       samara.AsyncProducer
	zookeeperNodes []string
)

type Request struct {
	Id                   int       `json:"id"`
	Type                 string    `json:"type"`
	Date                 time.Time `json:"date"`
	RejectReason         string    `json:"rejectReason"`
	SettlementAdjustment time.Time `json:"settlementAdjustment"`
	GrossValue           int       `json:"grossValue"`
	UserId               string    `json:"userId"`
	TechnologyType       int       `json:"tecnologyType"`
	StatusRequest        string    `json:"requestStatus"`
	Customer             string    `json:"customer"`
	ModCustomer          int       `json:"modCustomer"`
	BatchId              int       `json:"batchId"`
	MovementType         int       `json:"movementType"`
	RefundStatus         int       `json:"refundStatus"`
	AdjustmentReason     int       `json:"adjustmentReason"`
	AdjustmentComments   string    `json:"AdjustmentComments"`
	FileId               int       `json:"fieldId"`
	Product              int       `json:"product"`
	EntryType            string    `json:"entryType"`
	ProductCode          int       `json:"productCode"`

	CardAssociation     string `json:"cardAssociation"`
	FundingCurrencyCode string `json:"fundingCurrencyCode"`
}

func Init() {
	brokers = []string{"localhost:9092"}

	config = consumergroup.NewConfig()

	config.Offsets.Initial = samara.OffsetNewest

	config.ClientID = "AjusteGroup"

	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString("localhost:2181")

	go receive()
	log.Print("Kafka consumer iniciado")

}

func receive() {

	consumer, err := consumergroup.JoinConsumerGroup("ajustegroup", []string{"persistence"}, zookeeperNodes, config)

	if err != nil {
		log.Panic("Falha ao iniciar consumer", err)
	}

	log.Printf("Recebendo mensagens")

	var accumulator []Request
	for message := range consumer.Messages() {
		var req Request
		err := json.Unmarshal(message.Value, &req)
		if err != nil {
			log.Printf("Falha de parser na recepcao! %v", err)
		} else {
			accumulator = append(accumulator, req)
			if len(accumulator) > 3000 {
				go Persist(accumulator...)
				accumulator = []Request{}
			}
		}

	}

	log.Printf("Recebendo mensagens finalizado")
}
