package main

import "fmt"

func main() {
	fmt.Println(sum(10, 2, 5))
	fmt.Println(sum(111, 12, 35))
	fmt.Println(sum(13, 4, 88))
	fmt.Println(sum(-1, 2, 44))
	fmt.Println(sum(24, 7, 78))
}

func sum(numbers ...int) int {
	var out int

	for _, n := range numbers {
		out += n
	}
	return out
}
