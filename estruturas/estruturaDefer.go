package main

import "fmt"

func helloWord() {
	// mesmo declarando essa insrução no inicio da função ela é apenas executada quando a função termina
	defer fmt.Println("world")

	fmt.Println("hello")
}

func count() {
	// a resolução de vários DEFER segue como um LIFO
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	helloWord()
	count()
}
