package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
}

type Cow struct {
}

type Snake struct {
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func createAnimal(animalType string) Animal {
	switch animalType {
	case "bird":
		return Bird{}
	case "cow":
		return Cow{}
	case "snake":
		return Snake{}
	default:
		return nil
	}
}

func main() {
	animals := make(map[string]Animal)
	for true {
		fmt.Print("> ") // command name type/method
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		userInp := sc.Text()
		args := strings.Split(userInp, " ")

		if args[0] == "newanimal" {
			if _, ok := animals[args[1]]; ok {
				fmt.Println("An animal with such name already exists. Please use another one")
				continue
			}
			an := createAnimal(args[2])
			if an != nil {
				animals[args[1]] = an
				fmt.Println("Created it!")
			} else {
				fmt.Println("Unknown animal")
				continue
			}
		} else if args[0] == "query" {
			an, ok := animals[args[1]]
			if !ok {
				fmt.Println("There is no such animal")
				continue
			}
			switch args[2] {
			case "eat":
				an.Eat()
			case "move":
				an.Move()
			case "speak":
				an.Speak()
			default:
				fmt.Println("There is no such method")
				continue
			}
		} else {
			fmt.Println("Sorry, unknown command")
		}

	}
}
