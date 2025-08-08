package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(input string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // 함수 종료 후 실행됨
	length = len(input)
	uppercase = strings.ToUpper(input)
	return // length, uppercase 가 생략 가능
}

func main() {
	target := "abc"
	fmt.Println("before", target)
	_, changed := lenAndUpper(target)
	fmt.Println("after", changed)
}
