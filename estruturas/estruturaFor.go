package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sum := 0
	// for basico
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// omissão da estrutura inicializante e estrutura finalizadora, se comporta como o while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// loop infinito
	for {
		fmt.Println(rand.Intn(sum))
	}
}
