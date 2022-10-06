// Define a square struct with constructor and a couple of methods

package main

import (
	"fmt"
	"log"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length should be > 0")
	}
	s := Square{x, y, length}
	return &s, nil
}

func (s *Square) move(distX int, distY int) {
	s.X += distX
	s.Y += distY
}

func (s *Square) area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(10, 1, 5)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	} else {
		fmt.Println(s)
	}

	s.move(1, 5)
	fmt.Printf("%+v\n", s)
	fmt.Printf("Area: %v", s.area())
}
