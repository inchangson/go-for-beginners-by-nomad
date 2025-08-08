package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge > 18 {
		return true
	}
	fmt.Println("Are You Crazy?")
	return false
}

func main() {
	fmt.Println(canIDrink(17))
}
