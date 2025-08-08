package main

import "fmt"

func main() {
	names := [5]string{"tony", "jenn", "tim"}
	names[3] = "ken"
	names[4] = "austin"
	fmt.Println(names)
}
