package main

import "fmt"

func superAdd(numbers ...int) int {
	ret := 0
	for idx, number := range numbers {
		fmt.Printf("(%d, %d) is added\n", idx, number)
		ret += number
	}
	fmt.Println()

	cache := ret
	ret = 0
	for _, number := range numbers { // _ 를 사용하면 idx 무시 가능
		ret += number
	}
	fmt.Println()

	if ret != cache {
		fmt.Println("something goes wrong(1)")
	}

	ret = 0
	// index 통해 직접 접근 또한 가능
	for i := 0; i < len(numbers); i++ {
		ret += numbers[i]
	}

	if ret != cache {
		fmt.Println("something goes wrong(2)")
	}

	return ret
}

func main() {
	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}
