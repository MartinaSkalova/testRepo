package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("34 + 65 is %d \n", add(34, 65))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
