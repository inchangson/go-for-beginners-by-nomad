package main

import "fmt"

func main() {
	a := 2
	b := a
	fmt.Println("print memory address of a, b", &a, &b)
	fmt.Println("print values of a, b before change a", a, b)
	a = 10
	fmt.Println("print values of a, b after change a", a, b)

	c := &a
	fmt.Println("print memory addrss of a and value of c", &a, c)
	fmt.Println("print values of a, *c before change a", a, *c)
	a = 11
	fmt.Println("print values of a, *c after change a", a, *c)

	fmt.Println("print values of a, *c before change *c", a, *c)
	*c = 12
	fmt.Println("print values of a, *c after change *c", a, *c)

}
