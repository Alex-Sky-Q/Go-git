package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	sent := "hi.see you. "
	fmt.Println(CorrectSent(sent))
}

// Task 1. Write a function that:
// - converts first letter of a sentence to a capital letter
// - adds a dot at the end
// Assumptions:
//- "x.z" (no space after dot) - z is not a new sentence
//- "!", "?" at the end do not require additional dot

func CorrectSent(s string) (res string) {
	////var temp string
	//temp := strings.ToUpper(string(s[0]))
	//for i, l := range s[1:] {
	//	if string(l) == "." {
	//
	//	}
	//	temp += string(l)
	//}
	//return temp

	if strings.TrimSpace(s) == "" {
		return s
	}

	firstInd := 0
	for i, l := range s {
		if unicode.IsLetter(l) {
			res = strings.Replace(s, string(l), strings.ToUpper(string(l)), 1)
			firstInd = i
			break
		}
	}

	sentLoc, _ := regexp.Compile("[.?!]\\s+")
	loc := sentLoc.FindStringIndex(res[firstInd:])
	if len(loc) >= 2 {
		if loc[1] != len(res) {
			res = res[:loc[1]] + strings.ToUpper(string(res[loc[1]])) + res[loc[1]+1:]
		}
	}

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

	if !strings.ContainsAny(string(res[len(res)-1]), ".!?") {
		res += "."
	}
	return res
}
