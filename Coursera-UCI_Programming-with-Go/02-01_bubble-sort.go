package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Swap(sl []int, i int) {
	el0 := sl[i]
	sl[i] = sl[i+1]
	sl[i+1] = el0
}

func BubbleSort(sl []int) {
	sorted := false
	for !sorted {
		swapped := false
		for i := 0; i < len(sl)-1; i++ {
			if sl[i] > sl[i+1] {
				Swap(sl, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
	}
}

func main() {
	//a := []int{6, 1, 3, 2}
	var userSlice []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a set of ints (up to 10), separated by space")
	scanner.Scan()
	userString := scanner.Text()

	userStringSlice := strings.Split(userString, " ")
	if len(userStringSlice) > 10 {
		log.Fatalf("Sorry, there are more than 10 ints in a set")
	}
	
	for _, v := range userStringSlice {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Sorry, can't convert input to an int")
		}
		userSlice = append(userSlice, val)
	}

	BubbleSort(userSlice)
	fmt.Println(userSlice)
}
