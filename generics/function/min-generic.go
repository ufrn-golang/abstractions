// Example to demonstrate the implementation of a generic function
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 10, 2023

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Function min accepts any two values of a type parameter T that allows
// using inequality operators (namely integer, float-pointing, and string)
func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	a := 4
	b := 12
	fmt.Printf("min(%d, %d) = %d\n", a, b, min(a, b))

	x := "Z"
	y := "a"
	fmt.Printf("min(%s, %s) = %s\n", x, y, min(x, y))
}
