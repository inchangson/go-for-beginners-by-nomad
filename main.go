package main

import (
	"fmt"

	"github.com/inchangson/go-for-beginners-by-nomad/mydict"
)

func main() {
	fmt.Println("[START] Dictionary App")
	dictionary := mydict.Dictionary{}
	fmt.Println("1. Adding a word (hello, greeting)")
	err := dictionary.Add("hello", "greeting")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("2. Trying to Add existed word (hello)")
	err = dictionary.Add("hello", "greeting")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("3. Searching for a word existed(hello)")
	definition, error := dictionary.Search("hello")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

	fmt.Println("4. Searching for a word that are not existed(bye)")
	definition, error = dictionary.Search("bye")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

	fmt.Println("5. Update for a word (hello, greeting) to (hello, again)")
	err = dictionary.Update("hello", "again")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dictionary.Search("hello"))
	}

	fmt.Println("6. Delete for a word (hello, greeting)")
	err = dictionary.Delete("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted successfully, searching for the word again:")
		fmt.Println(dictionary.Search("hello"))
	}

	fmt.Println("[END] Dictionary App")
}
