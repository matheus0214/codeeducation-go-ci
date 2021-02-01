package main

import (
	"fmt"
)

func main() {
	res := Sum(1, 2)

	fmt.Printf("%d\n", res)
}

// Sum Get two numbers an return sum
func Sum(a int, b int) int {
	return a + b
}
