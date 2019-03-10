package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[string]string)
	c := make(chan string)

	mutex := sync.Mutex{}
	go func() {
		mutex.Lock()
		m["a"] = "abcd"
		mutex.Unlock()
		c <- "adsds"
	}()
	mutex.Lock()
	m["b"] = "fsds"
	mutex.Unlock()

	<-c

	fmt.Printf("%v", m)
}
