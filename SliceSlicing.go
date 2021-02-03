package main

import "fmt"

func SliceSlicing() {
	s := []int{2, 3, 5, 7, 11, 13}
	slice(s)

	// Slice the slice to give it zero length
	s = s[:0]
	slice(s)

	// Extend its length
	s = s[:4]
	slice(s)

	// Drop its first two values
	s = s[2:]
	slice(s)

}

// simple print function
func slice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
