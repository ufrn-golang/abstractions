// Example to demonstrate the implementation of methods
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 10, 2023

package main

import "fmt"

type Rectangle struct {
	width  float32
	height float32
}

func (r Rectangle) area() float32 {
	return r.width * r.height
}

func (r Rectangle) printArea() {
	fmt.Println("Rectangle's area:", r.area())
}

func main() {
	r1 := Rectangle{3, 4}
	r1.printArea()

	r2 := Rectangle{4, 4}
	r2.printArea()
}
