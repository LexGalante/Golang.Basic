//nome do pacote
package main

//importes de outros pacotes, ver GOPATH
import (
	"fmt"
	"math/rand"
)

//função inicializadora do programa
func main() {
	fmt.Println("Número aleatorio: ", rand.Intn(100))
}
