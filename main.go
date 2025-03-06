package main

import "fmt"

func main() {
	fmt.Println(divide(10, 5))
}

func divide(a, b int) int {
	res := a / b
	return res
}
