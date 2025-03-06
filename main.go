package main

import "fmt"

func main() {
	fmt.Println(twoParams(10, 5))
}

func twoParams(a, b int) (n1, n2 int) {
	n1 = a + b
	n2 = a * b
	return n1, n2
}
