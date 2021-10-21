package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // ponteiro para i
	fmt.Println(*p) // le o valor de i atraves do ponteiro
	*p = 21         // define o valor de i atraves do ponteiro
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println("&i = ", &i)
}
