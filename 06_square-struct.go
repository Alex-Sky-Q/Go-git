// Define a square struct with constructor and a couple of methods

package main

import "fmt"

type Square struct {
	x      int
	y      int
	length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length should be >= 0")
	}
	s := Square{x, y, length}
	return &s, nil
}

func (s *Square) move(distX int, distY int) {
	s.x += distX
	s.y += distY
}

func (s *Square) area() (area int) {
	return s.length * s.length
}

func main() {
	s, err := NewSquare(10, 1, 5)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(s)
	}

	s.move(1, 5)
	fmt.Println(s)
	fmt.Printf("Area: %v", s.area())
}
