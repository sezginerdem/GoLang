package main

import "fmt"

func main() {

	name := "Sezgin"

	age := 40

	isMarried := true

	name = "Erdem"

	isMarried = false

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
}
