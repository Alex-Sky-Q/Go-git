package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func splitSliceIndexes(slice *[]int, splitAmount int) []int {
	splitFactor := len(*slice) / splitAmount
	indexes := []int{splitFactor}
	for i := 0; i < 2; i++ {
		indexes = append(indexes, indexes[i]+splitFactor)
	}
	return indexes
}

func sortSlice(slice *[]int, group *sync.WaitGroup) {
	fmt.Printf("Going to sort a slice: %v\n", *slice)
	sort.Ints(*slice)
	fmt.Printf("Sorted slice: %v\n", *slice)
	group.Done()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a series of ints at least 4 ints long")
	sc.Scan()
	userString := sc.Text()

	userStringsSlice := strings.Split(userString, " ")
	if len(userStringsSlice) < 4 {
		log.Fatalf("There are less than 4 ints")
	}

	var userNumSlice []int

	for _, s := range userStringsSlice {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Cannot convert '%s' to int", s)
		}
		userNumSlice = append(userNumSlice, num)
	}

	indexes := splitSliceIndexes(&userNumSlice, 4)

	slice1 := userNumSlice[:indexes[0]]
	slice2 := userNumSlice[indexes[0]:indexes[1]]
	slice3 := userNumSlice[indexes[1]:indexes[2]]
	slice4 := userNumSlice[indexes[2]:]

	var wg sync.WaitGroup
	wg.Add(4)
	go sortSlice(&slice1, &wg)
	go sortSlice(&slice2, &wg)
	go sortSlice(&slice3, &wg)
	go sortSlice(&slice4, &wg)
	wg.Wait()

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	slice12 := append(slice1, slice2...)
	slice34 := append(slice3, slice4...)
	finSlice := append(slice12, slice34...)

	// It's very strange that according to task we need to sort a final slice after splitting and sorting sub-slices.
	// It's inefficient, but is done exactly according to task requirements
	sort.Ints(finSlice)
	fmt.Println(finSlice)
}
