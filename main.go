package main

import "fmt"

func main() {
	const name string = "nicolas"
	// name = "Lynn" // cannot change const
	fmt.Println(name)

	var age int = 30
	fmt.Println("before", age)
	age = 31 // can change var
	fmt.Println("after", age)

	weight := 70.5 // support type inference only if ⭐️it is in the function block⭐️
	// => var weight float32 = 70.5
	fmt.Println(weight, "kg")
}
