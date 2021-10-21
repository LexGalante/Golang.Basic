package main

import "fmt"

// go não tem classes apenas estruturas
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Println(v.X)
	fmt.Println(v.Y)

	v.X = v.Y
	fmt.Println(v.Y)

	// ponteiro para v
	p := &v

	fmt.Println(v.X)
	fmt.Println(v.Y)

	p.X = p.X * p.Y

	fmt.Println(v.X)
	fmt.Println(v.Y)

	// atribuições de campos em structs
	v1 := Vertex{1, 2} // define os valores pelas ordens que estão declarados na struct
	v2 := Vertex{X: 1} // utiliza o nome do atributo para atribuir o valor
	v3 := Vertex{}     // X:0 and Y:0
	p = &Vertex{1, 2}  // ponteiro para *Vertex

	fmt.Println(v1, p, v2, v3)
}
