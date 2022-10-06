/*
An "even ended number" is a number whose first and last digits are the same.
Your mission is to count how many "even ended numbers" are there that are a multiplication of two 4-digit numbers
*/

package main

import "fmt"

func main() {
	cnt := 0
	for i := 1000; i <= 9999; i++ {
		for n := i; n <= 9999; n++ {
			mp := i * n
			mpStr := fmt.Sprintf("%d", mp)
			if mpStr[0] == mpStr[len(mpStr)-1] {
				cnt += 1
			}
		}
	}
	fmt.Printf("There are %v even ended numbers", cnt)
}
