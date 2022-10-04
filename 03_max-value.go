//Calculate max value in a slice {16, 8, 42, 4, 23, 15}

package main

import "fmt"

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	maxVal := nums[0]
	for _, num := range nums[1:] {
		if num > maxVal {
			maxVal = num
		}
	}
	fmt.Printf("The max value is %v", maxVal)
}
