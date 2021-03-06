////////////////////////////////////////////////////////////////////////////////
//	requestLoader.go  -  May/06/2017  -  aldebaran perseke
//
//	Entry point of the request loader component
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"os"
	"requestLoader"
)

////////////////////////////////////////////////////////////////////////////////
//	Main function of the request loader component
////////////////////////////////////////////////////////////////////////////////

func main() {

	var loadDirectory string
	var ignoreHeader bool
	var kafkaTopic string
	var verbose bool

	//	parse command line arguments
	flag.StringVar(&loadDirectory, "loadDirectory", "input", "set the directory where get load files")
	flag.BoolVar(&ignoreHeader, "ignoreHeader", true, "ignore the header line from the load files")
	flag.StringVar(&kafkaTopic, "kafkaTopic", "request", "set the name of Kafka topic to produce messages to")
	flag.BoolVar(&verbose, "verbose", false, "print a detailed trace execution")

	flag.Parse()

	if 0 < len(flag.Args()) {
		fmt.Fprintf(os.Stderr, "invalid argument: %s", flag.Args())
		panic(-1)
	}

	//	validate the arguments
	_, err := os.Stat(loadDirectory)
	if nil != err {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "[error] invalid directory name %s\n", loadDirectory)
		} else {
			fmt.Fprintf(os.Stderr, "[error] can't get info about directory name %s\n", loadDirectory)
		}
	}

	//	splash screen
	fmt.Printf(">>>>> Request Loader %s\n\n", loadDirectory)

	//	initialize Kafka connection
	requestLoader.Init([]string{"localhost:9092"})

	//	polling the load directory for request files
	requestLoader.SetIgnoreHeader(ignoreHeader)
	requestLoader.SetVerbose(verbose)
	requestLoader.SetTopicName(kafkaTopic)

	requestLoader.LoadDirectoryPolling(loadDirectory)
}
