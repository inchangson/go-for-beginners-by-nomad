package main

import (
	"fmt"

	"github.com/inchangson/go-for-beginners-by-nomad/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "world"
	definition, error := dictionary.Search("hello")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

	definition, error = dictionary.Search("bye")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}
}
