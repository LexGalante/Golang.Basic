package main

import (
	"fmt"
	"strings"
)

func main() {
	// função make builtin cria os slices
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// adionando elementos a um slice
	adionandoElementosAoSlice()

	iterandoSlice()
}

func adionandoElementosAoSlice() {
	var s []int
	fmt.Println(s)

	// append works on nil slices.
	s = append(s, 0)
	fmt.Println(s)

	// The slice grows as needed.
	s = append(s, 1)
	fmt.Println(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}

func iterandoSlice() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
