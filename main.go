package main

import "fmt"

func main() {
	fmt.Println(sum(10)(20))
}

func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
