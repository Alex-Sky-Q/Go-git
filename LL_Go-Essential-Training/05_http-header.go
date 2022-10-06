// Write a function that returns Content-Type header

package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://google.c"
	cTypeHeader, err := getCType(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cTypeHeader)
	}
}

func getCType(url string) (head []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//defer func() {
	//	err = resp.Body.Close()
	//}()
	head = resp.Header["Content-Type"]
	return head, err
}
