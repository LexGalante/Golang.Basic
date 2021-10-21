package main

import "fmt"

//multideclaração
var python, csharp, java string = "python", "c#", "java"

func main() {
	// atribuição curta
	c, golang := "c", "golang"

	fmt.Println(python, csharp, java, c, golang)
}
