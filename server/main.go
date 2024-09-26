package main

import "sync"

func main() {
	// Open peer-discovery mode
	go RespondToPeerDiscovery()

	// Open file-receiving mode
	go ReceiveFile()

	// Wait until the program is killed
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
