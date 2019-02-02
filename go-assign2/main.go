package main

import "fmt"

type shape interface {
	Area() float64
}

type square struct {
	side float64
}

type triangle struct {
	base, height float64
}

func main() {
	sq := square{3}
	tr := triangle{3, 2}

	printArea(sq)
	printArea(tr)

}

func (s square) Area() float64 {
	return s.side * s.side
}

func (t triangle) Area() float64 {
	return t.base * t.height * 0.5
}

func printArea(s shape) {
	fmt.Println(fmt.Sprintf("%f", s.Area()))
}
