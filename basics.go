package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sent := "hello world. with pleasure"
	fmt.Println(CorrectSent(sent))
}

// Task 1. Write a function that:
// - converts first letter of a sentence to a capital letter
// - adds a dot at the end
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
	var temp []string
	strSplit := strings.Split(s, ". ")

	for _, t := range strSplit {
		for _, l := range t {
			if unicode.IsLetter(l) {
				t = strings.Replace(t, string(l), strings.ToUpper(string(l)), 1)
				break
			}
		}
		temp = append(temp, t)
	}

	res = strings.Join(temp, ". ")

	if string(res[len(res)-1]) != "." {
		res += "."
	}
	return res
}
