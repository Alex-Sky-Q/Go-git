// Write a function that returns Content-Type header

package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://google.com"
	cTypeHeader, err := getCType(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cTypeHeader)
	}
}

func getCType(url string) (head []string, err error) {
	resp, err := http.Get(url)
	//defer func(resp *http.Response) {
	//	err = resp.Body.Close()
	//	if err != nil {
	//		fmt.Println("test")
	//	}
	//}(resp)
	if err != nil {
		return
	} else {
		defer resp.Body.Close()
		head = resp.Header["Content-Type"]
		return head, err
	}
}
