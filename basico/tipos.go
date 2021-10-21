package main

import (
	"fmt"
	"math/cmplx"
)

var (
	boleano        bool       = false
	numero         uint64     = 1<<64 - 1
	numeroComplexo complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", boleano, boleano)
	fmt.Printf("Type: %T Value: %v\n", numero, numero)
	fmt.Printf("Type: %T Value: %v\n", numeroComplexo, numeroComplexo)
	// valores default
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
