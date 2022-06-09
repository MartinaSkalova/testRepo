package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Now you are %d people's problem.\n", math.MaxInt)
	fmt.Printf("34 + 65 is %d \n", add(34, 65))
}
