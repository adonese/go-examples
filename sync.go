package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	m map[string]int
	sync.Mutex
}

func (s *SafeMap) Create() {
	m := make(map[string]int)
	s.m = m
}
func (s *SafeMap) Inc(key string) {
	s.Lock()
	s.m[key]++
	s.Unlock()
}

func main() {

	var s SafeMap
	s.Create()
	key := "foo"

	for i := 0; i <= 20; i++ {
		go s.Inc(key)
	}
	s.Lock()
	fmt.Printf("The final key val is: %v", s.m[key])
	s.Unlock()
}
