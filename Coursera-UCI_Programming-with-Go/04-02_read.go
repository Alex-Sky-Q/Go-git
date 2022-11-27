package main

import (
	"unicode/utf8"
	"strings"
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	type name struct {
		fname string
		lname string
	}
	var fileName string
	var names []name

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please, enter a text file name")
	scanner.Scan()
	fileName = scanner.Text()
	
	f, err := os.Open(path.Join(fileName))
	if err != nil {fmt.Println("Error during file opening")}
	defer f.Close()

	fScanner := bufio.NewScanner(f)
	for fScanner.Scan() {
		line := strings.Split(fScanner.Text(), " ")
		if utf8.RuneCountInString(line[0]) > 20 || utf8.RuneCountInString(line[1]) > 20 {
			log.Fatal("First and Last name should be less then 20 chars long")
		}
		var tmpName name
		tmpName.fname = line[0]
		tmpName.lname = line[1]
		names = append(names, tmpName)	
	}
	
	for _, v := range names {
		fmt.Printf("%s %s \n", v.fname, v.lname)
	}
}