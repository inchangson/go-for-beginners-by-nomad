package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(input string) (int, string) {
	return len(input), strings.ToUpper(input)
}

func main() {
	target := "abc"
	fmt.Println("before", target)
	length, changed := lenAndUpper(target)
	fmt.Println("after", length, changed)
}
