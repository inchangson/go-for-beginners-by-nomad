package main

import (
	"fmt"

	"github.com/inchangson/go-for-beginners-by-nomad/something"
)

func main() {
	fmt.Println("Hello world!!!")
	something.SayHello()
	// something.sayBye() // cannot use as it is private
}
