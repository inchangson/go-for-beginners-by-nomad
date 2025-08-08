package main

import "fmt"

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age < 50:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(canIDrink(17))
	fmt.Println(canIDrink(30))
	fmt.Println(canIDrink(90))
}
