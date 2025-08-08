package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchiRamen", "chicken"}
	nico1 := person{"nico", 28, favFood}
	fmt.Println(nico1)

	nico2 := person{name: "nico", age: 28, favFood: favFood}
	fmt.Println(nico2)

}
