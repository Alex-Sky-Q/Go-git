package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	userSlice := make([]int, 0, 3)
	for true {
		fmt.Println("Please, enter a number, or 'X' to quit")
		var userInp string
		_, err := fmt.Scan(&userInp)
		if err != nil {
			fmt.Println("Sorry, an error occured")
		}

		if strings.ToUpper(userInp) == "X" {
			break
		}

		userInt, err := strconv.Atoi(userInp)
		if err != nil {
			fmt.Println("Sorry, there is an error")
		}

		userSlice = append(userSlice, userInt)
	}

	sort.Ints(userSlice)
	fmt.Println(userSlice)
}
