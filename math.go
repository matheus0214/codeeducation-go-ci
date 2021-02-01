package main

import (
	"fmt"
)

func main() {
	res := sum(1, 2)

	fmt.Printf("%d\n", res)
}

// Sum Get two numbers an return sum
func sum(a int, b int) int {
	return a + b
}
