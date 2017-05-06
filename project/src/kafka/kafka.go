package kafka

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	samara "gopkg.in/Shopify/sarama.v1"
)

var (
	channel  = make(chan string, 50000)
	process  Process
	counter  = int32(0)
	brokers  []string
	config   *samara.Config
	producer samara.AsyncProducer
)

type Process func(string)

func Init(proc Process, executors int, broker []string) {
	process = proc
	brokers = broker

	config = samara.NewConfig()
	config.ClientID = "ConsumerRaw"
	config.Consumer.Return.Errors = true
	config.Consumer.Fetch.Min = 1
	config.Producer.Flush.MaxMessages = 100000

	producer, _ = samara.NewAsyncProducer(brokers, config)

	for i := 0; i < executors; i++ {
		go executor(i)
	}

	go monitor()

}

func monitor() {
	for {
		time.Sleep(1 * time.Second)
		log.Printf("Received: %d", counter)
	}
}

func executor(id int) {
	log.Printf("Executor iniciando...")
	for payload := range channel {
		if strings.TrimSpace(payload) != "" {
			atomic.AddInt32(&counter, 1)
			process(payload)
		}
	}
}

func Send(destination string, data interface{}, id int) {
	payload, _ := json.Marshal(data)
	producer.Input() <- &samara.ProducerMessage{
		Topic: destination,
		Value: samara.ByteEncoder(payload),
		Key:   samara.StringEncoder(strconv.Itoa(id)),
	}
}

func Receive(topic string, offset int64) {
	master, err := samara.NewClient(brokers, config)
	if err != nil {
		log.Panicf("Falha ao estabelecer conexao %v", err)
	}

	consumerClient, err := samara.NewConsumerFromClient(master)

	if err != nil {
		log.Fatalf("Erro ao iniciar consumer client: %v", err)
	}
	partitions, _ := consumerClient.Partitions(topic)

	for _, data := range partitions {
		go consume(topic, data, 0)
	}

}

func consume(topic string, partition int32, offset int64) {
	master, err := samara.NewClient(brokers, config)
	consumerClient, err := samara.NewConsumerFromClient(master)
	consumer, err := consumerClient.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Fatalf("Erro ao iniciar consumer: %v", err)
	}

	log.Printf("Consumer[%v] Listening ...", partition)

	for {
		select {
		case message := <-consumer.Messages():
			channel <- string(message.Value)
		case err := <-consumer.Errors():
			log.Printf("Error: %v", err)
		}
	}
}
