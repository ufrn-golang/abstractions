// Example to demonstrate the implementation of a generic struct
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 10, 2023

package main

import "fmt"

// K is of int or string type
// V is of int, float64, or string type
// IndexTable associates a key of type K (int or string) to a slice
// of values of type V (int, float64 or string)
type IndexTable[K int|string, V int|float64|string] struct {
    key K
    values []V
}

func main() {
    // As the compiler fails to make the type inference in this case, 
    // it is required explicity adding the type parameters when
    // instantiating the struct
    table := IndexTable[string, int]{"1", []int{1, 2, 3, 4, 5}}
    fmt.Printf("%v -> %v\n", table.key, table.values)
}
