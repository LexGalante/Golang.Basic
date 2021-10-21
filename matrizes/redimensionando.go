package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// valor default de um slice é nil
	var ss []int
	printSlice(ss)
}

func printSlice(s []int) {
	if s == nil {
		fmt.Println("nil!")
	} else {
		fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	}
}
