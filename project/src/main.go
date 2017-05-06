package main

import (
	"app"
	"kafka"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {

	kafka.Init(app.Process, 20, []string{"localhost:9092"})

	wg.Add(1)

	go kafka.Receive("request", int64(0))

	wg.Wait()
}
