package main

import "fmt"

func canIDrink(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge < 50:
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
