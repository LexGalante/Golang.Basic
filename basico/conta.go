package main

import (
	"fmt"
	"math/rand"
)

// retorno de dois valores
func geraInteiroAleatorio(limite int) (int, int) {
	return rand.Intn(limite), rand.Intn(limite)
}

// retornos nomeados
func elevaInteiroAleatorio(a int, b int) (x, y int) {
	x = a ^ b
	y = b * a
	return
}

// tipagem de parametros
func soma(a int, b int) int {
	return a + b
}

// tipagem apenas no ultimo parametro
func subtracao(a, b int) int {
	return a - b
}

func main() {
	a, b := geraInteiroAleatorio(100)
	x, y := elevaInteiroAleatorio(a, b)

	fmt.Println("A: ", a)
	fmt.Println("B: ", b)

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)

	fmt.Println("SOMA: ", soma(a, b))
	fmt.Println("SUBTRACAO: ", subtracao(a, b))

}
