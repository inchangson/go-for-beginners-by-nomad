package main

import (
	"fmt"

	"github.com/inchangson/go-for-beginners-by-nomad/mydict"
)

func main() {
	fmt.Println("[START] Dictionary App")
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("hello", "greeting")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("trying to add a word that already exists")
	err = dictionary.Add("hello", "greeting")
	if err != nil {
		fmt.Println(err)
	}

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
	fmt.Println("[END] Dictionary App")
}
