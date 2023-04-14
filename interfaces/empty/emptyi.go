// Example to demonstrate the use of empty interfaces.
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 10, 2023

package main

import "fmt"

func main() {
	var i interface{}

	// i is of type int...
	i = 42
	fmt.Printf("Value %v\n", i)

	// ...now i is of type string
	i = "Hello"
	fmt.Println("Value:", i)

	switch i.(type) {
	case int:
		fmt.Println("The variable is an integer")
	default:
		fmt.Println("The variable is not an integer")
	}
}
