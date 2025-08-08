package main

import "fmt"

func canIDrink(age int) bool {
	if age > 18 {
		return true
	}
	fmt.Println("Are You Crazy?")
	return false
}

func main() {
	fmt.Println(canIDrink(18))
}
