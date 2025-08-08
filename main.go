package main

import (
	"fmt"
	"strconv"
)

func multiply(a int, b int) int {
	return a * b
}

func add(a, b, c int) int { // you can write type only once if it is same
	return a + b + c
}

func myPrint(input string) { // void function has no return type
	fmt.Println("===RESULT===\n", input, "\n============")
}

func main() {
	myPrint(strconv.Itoa(multiply(2, 2)))
	myPrint(strconv.Itoa(add(2, 2, 2)))
}
