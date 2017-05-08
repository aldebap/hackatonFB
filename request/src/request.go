////////////////////////////////////////////////////////////////////////////////
//	request.go  -  May/07/2017  -  aldebaran perseke
//
//	Request RESTFul service
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"requestLoader"

	"github.com/gorilla/mux"
)

////////////////////////////////////////////////////////////////////////////////
//	RESTFul service to add a new request to Kafka topic
////////////////////////////////////////////////////////////////////////////////

func insertRequest(_httpResponse http.ResponseWriter, _httpRequest *http.Request) {

	var request requestLoader.Request

	_ = json.NewDecoder(_httpRequest.Body).Decode(&request)

	requestLoader.SendTopic(request)

	_httpResponse.WriteHeader(http.StatusCreated)
}

////////////////////////////////////////////////////////////////////////////////
//	Start the RESTFul server
////////////////////////////////////////////////////////////////////////////////

func main() {

	var kafkaTopic string
	var verbose bool

	//	parse command line arguments
	flag.StringVar(&kafkaTopic, "kafkaTopic", "request", "set the name of Kafka topic to produce messages to")
	flag.BoolVar(&verbose, "verbose", false, "print a detailed trace execution")

	flag.Parse()

	if 0 < len(flag.Args()) {
		fmt.Fprintf(os.Stderr, "invalid argument: %s", flag.Args())
		panic(-1)
	}

	//	splash screen
	fmt.Printf(">>>>> Request RESTFul Service\n\n")

	//	initialize Kafka connection
	requestLoader.Init([]string{"localhost:9092"})

	//	start the Web Server
	requestLoader.SetVerbose(verbose)
	requestLoader.SetTopicName(kafkaTopic)

	router := mux.NewRouter()

	router.HandleFunc("/request", insertRequest).Methods("POST")

	log.Panic(http.ListenAndServe(":8080", router))
}
