package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	sent := "-"
	fmt.Println(CorrectSent(sent))
}

// Task 1. Write a function that:
// - converts first letter of a sentence to a capital letter
// - adds a dot to the end
// Assumptions:
// ". ", "! " or "? " - are the only valid sentences' separators. If smth like ".- " should also be checked,
// it can be done by updating regex
//- "x.z" (no space after dot in the middle of a string) - z is not a new sentence
//- "!", "?" at the end do not require additional dot

func CorrectSent(s string) (res string) {

	if strings.TrimSpace(s) == "" {
		return s
	}

	res = s

	// Checking the beginning of a string
	firstInd := 0
	for i, l := range s {
		if unicode.IsLetter(l) {
			res = strings.Replace(s, string(l), strings.ToUpper(string(l)), 1)
			firstInd = i
			break
		}
	}

	sentLoc, _ := regexp.Compile("[.?!]\\s+")
	loc := sentLoc.FindAllStringIndex(res[firstInd:], -1)
	for _, l := range loc {
		if l[1] != len(res) {
			res = res[:l[1]] + strings.ToUpper(string(res[l[1]])) + res[l[1]+1:]
		}
	}

	// First solution, does not work in case of different endings of the sentence ("!", "?")
	//var temp []string
	//strSplit := strings.Split(s, ". ")
	//for _, t := range strSplit {
	//	for _, l := range t {
	//		if unicode.IsLetter(l) {
	//			t = strings.Replace(t, string(l), strings.ToUpper(string(l)), 1)
	//			break
	//		}
	//	}
	//	temp = append(temp, t)
	//}
	//res = strings.Join(temp, ". ")

	// Checking the end of a string
	if !strings.ContainsAny(string(res[len(res)-1]), ".!?") {
		res += "."
	}

	return res
}
