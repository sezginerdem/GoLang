package main

import "fmt"

var packVar = "Package Scope"

func main() {

	if true {
		var bloackVar = "Block Scope"
		fmt.Println(bloackVar)
	}

	var funcVar = "Func scope"

	fmt.Println(funcVar)
	fmt.Println(packVar)
}
