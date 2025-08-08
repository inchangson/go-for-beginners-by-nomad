package main

import (
	"fmt"
	"time"
)

func isWeekend() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend !!!")
	default:
		fmt.Println("I'm sorry")
	}
}

func main() {
	isWeekend()
}
