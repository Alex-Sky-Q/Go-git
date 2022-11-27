package main

import (
	"strings"
	"regexp"
	"fmt"
)

// func CheckString(userStr string) string {
func CheckString() {
	var userString string
	fmt.Println("Enter a string:")
	_, err := fmt.Scan(&userString)
	if err != nil {
		fmt.Println("Sorry, an error occured")
	}
	// userString := userStr
	re := regexp.MustCompile(`^[i].*a+.*[n]$`)

	if re.MatchString(strings.ToLower(userString)) {
		fmt.Println("Found!")
		// return "Found!"
	} else {
		fmt.Println("Not Found!")
		// return "Not Found!"
	}
}

func main() {
	CheckString()
}
