package main

import (
	"fmt"
	"math"
)

// acceleration a, initial velocity vo, initial displacement so
func GenDisplaceFn(a, Vo, So float64) func(t float64) float64 {
	return func(t float64) float64 {
		s := (a*math.Pow(t, 2))/2 + Vo*t + So
		return s
	}
}

func main() {
	var acceleration float64
	fmt.Println("Please, enter acceleration:")
	_, err := fmt.Scan(&acceleration)
	if err != nil {
		fmt.Println("Sorry, an error occurred")
	}

	var velocity float64
	fmt.Println("Please, enter velocity:")
	_, err = fmt.Scan(&velocity)
	if err != nil {
		fmt.Println("Sorry, an error occurred")
	}

	var initDisplacement float64
	fmt.Println("Please, enter initial displacement:")
	_, err = fmt.Scan(&initDisplacement)
	if err != nil {
		fmt.Println("Sorry, an error occurred")
	}

	fn := GenDisplaceFn(acceleration, velocity, initDisplacement)

	var moveTime float64
	fmt.Println("Please, enter time:")
	_, err = fmt.Scan(&moveTime)
	if err != nil {
		fmt.Println("Sorry, an error occurred")
	}

	fmt.Printf("The displacement is: %v\n", fn(moveTime))
}
