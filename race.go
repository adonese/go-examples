package main

import "fmt"

func main() {
	m := make(map[string]string)
	c := make(chan string)

	go func() {
		m["a"] = "abcd"
		c <- "adsds"
	}()
	m["b"] = <-c
	//<-c
	fmt.Printf("%v", m)
}
