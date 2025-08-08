package main

import "fmt"

func main() {
	names := []string{"tony", "jenn", "tim"}
	names = append(names, "ken")
	names = append(names, "austin")
	fmt.Println(names)
}
