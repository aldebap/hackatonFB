package kafka

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/wvanbergen/kafka/consumergroup"
	"github.com/wvanbergen/kazoo-go"

	samara "gopkg.in/Shopify/sarama.v1"
)

var (
	channel        = make(chan string, 50000)
	process        Process
	counter        = int32(0)
	brokers        []string
	config         *consumergroup.Config
	producer       samara.AsyncProducer
	zookeeperNodes []string
)

type Process func(string)

func Init(proc Process, executors int, broker []string) {
	process = proc
	brokers = broker

	config = consumergroup.NewConfig()

	config.Offsets.Initial = samara.OffsetNewest

	configProducer := samara.NewConfig()
	configProducer.ClientID = "ConsumerRaw"
	configProducer.Producer.Flush.MaxMessages = 100000

	zookeeperNodes, config.Zookeeper.Chroot = kazoo.ParseConnectionString("localhost:2181")

	producer, _ = samara.NewAsyncProducer(brokers, configProducer)

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

//https://github.com/wvanbergen/kafka/blob/master/consumergroup/consumer_group.go
func Send(destination string, data interface{}, id int) {
	payload, _ := json.Marshal(data)
	producer.Input() <- &samara.ProducerMessage{
		Topic: destination,
		Value: samara.ByteEncoder(payload),
		Key:   samara.StringEncoder(strconv.Itoa(id)),
	}
}

func Receive(topic string, offset int64) {

	consumer, err := consumergroup.JoinConsumerGroup("groupservice", []string{"request"}, zookeeperNodes, config)

	if err != nil {
		log.Panic("Falha ao iniciar consumer", err)
	}

	for message := range consumer.Messages() {
		channel <- string(message.Value)
	}

}
