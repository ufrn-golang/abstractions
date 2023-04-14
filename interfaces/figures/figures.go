// Example to demonstrate the implementation of an interface and
// struct types satisfying it through methods
//
// Author: Everton Cavalcante (everton.cavalcante@ufrn.br)
// Date: April 10, 2023

package main

import "fmt"

type Figure interface {
	area() float32
}

type Rectangle struct {
	width, height float32
}

type Square struct {
	side float32
}

func (r Rectangle) area() float32 {
	return r.width * r.height
}

func (q Square) area() float32 {
	return q.side * q.side
}

func main() {
	rect := Rectangle{3, 4}
	sq := Square{2}
	fmt.Println("Rectangle's area:", rect.area())
	fmt.Println("Square's area:", sq.area())
}
