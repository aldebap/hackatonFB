package main

import (
	"service"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	service.Init()

	wg.Wait()

}
