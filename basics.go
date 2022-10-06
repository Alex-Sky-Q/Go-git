package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	sent := "hello world. with pleasure"
	//fmt.Println(string(sent[1]))
	fmt.Println(correctSent(sent))
}

func correctSent(s string) (res string) {
	////var temp string
	//temp := strings.ToUpper(string(s[0]))
	//for i, l := range s[1:] {
	//	if string(l) == "." {
	//
	//	}
	//	temp += string(l)
	//}
	//return temp
	var temp []string
	strSplit := strings.Split(s, ". ")

	for _, t := range strSplit {
		for _, l := range t {
			if unicode.IsLetter(l) {
				t = strings.Replace(t, string(l), strings.ToUpper(string(l)), 1)
				break
			}
		}
		//t = strings.ToUpper(string(t[0])) + t[1:]
		temp = append(temp, t)
	}

	res = strings.Join(temp, ". ")

	if string(res[len(res)-1]) != "." {
		res += "."
	}
	return res
}
