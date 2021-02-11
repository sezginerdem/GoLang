package main

import "fmt"

func main() {

	/* var (
		name      = "Sezgin"
		age       = 40
		isMarried = true

		weight float32 = 72.5
		height int     = 172
	) */

	/* var name, age, isMarried, weight, height = "Sezgin", 40, true, 72.5, 172 */
	name, age, isMarried, weight, height := "Sezgin", 40, true, 72.5, 172

	var name string

	fmt.Print(name) // string --> "", numeric --> 0. boolean --> False

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
	fmt.Println(height)

}
