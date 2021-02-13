/* package main

import "fmt"

func main() {
	/*
		var studentName string = "John Doe"
		var grade int = 77
		var isPassed bool = true */

/* 	var studentName = "John Doe"
   	var grade = 77
   	var isPassed = true */

/* 	studentName := "John Doe"
	grade := 77
	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
} */

// yukaridaki degiskenleri tek satirda yazmak

/* package main

import "fmt"

func main() {

	/* 	studentName := "John Doe"
	   	grade := 77
	   	isPassed := true
*/
/* var studentName, grade, isPassed = "John Doe", 77, true */
/* 	studentName, grade, isPassed := "John Doe", 77, true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
} */

// decleration, assign, initialization. initial value

/* package main

import "fmt"

func main() {

	var studentName string = "John Doe"
	studentName = "Sezgin Erdem"

	fmt.Println(studentName)
} */

//Statically(bir degiskeni bir kez verdi isen onu baska bir yerde degistiremezsin. veri tipini bildigi icin hafizayi daha ekonomik kullanir) ve dynamically(veri tipi degisebiliyor pythondaki gibi. kod yazimi daha kolay esneklik sagliyor) kavramlari

// := ve = arasindaki fark

package main

import "fmt"

func main() {

	/* 	var studentName string = "John Doe"
	   	studentName string = "Sezgin Erdem" */

	studentName := "john doe"
	studentName = "Sezgin erdem"
	fmt.Println(studentName)
}
