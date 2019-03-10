// SliceBytes really does nothing.
package main

import "fmt"

type SliceBytes []byte

func main() {
	var s SliceBytes
	b := []byte("abc")
	Append(s, b)
	fmt.Printf("The non-pointer returns: %v\n", string(s))
	s.Append(b)
	fmt.Printf("The pointer returns: %v\n", string(s))

}

func Append(s SliceBytes, b []byte)  {
	s = append(s, b...)
}

func (s *SliceBytes) Append(b []byte) {
	*s = append(*s, b...)
}
