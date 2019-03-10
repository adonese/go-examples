package main

import "fmt"

func main() {
	a := 1
	zeroVal(a)
	fmt.Printf("Pass by value %v\n", a)
	zeroPtr(&a)
	fmt.Printf("Pass by reference %v\n", a)
	fmt.Printf("The address of i is: %v", &a)

}

func zeroVal(a int) {
	a = 100000
}

func zeroPtr(a *int) {
	*a = 1000000
}
