package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func makeJson() {
	var name string
	var address string
	person := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Please, enter a name")
	scanner.Scan()
	name = scanner.Text()
	person["name"] = name

	fmt.Println("Please, enter an address")
	scanner.Scan()
	address = scanner.Text()
	person["address"] = address

	barr, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Sorry, there is an error")
	}

	fmt.Println(string(barr))
}
