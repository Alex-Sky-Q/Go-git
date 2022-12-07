package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (an Animal) Eat() {
	fmt.Println(an.food)
}

func (an Animal) Move() {
	fmt.Println(an.locomotion)
}

func (an Animal) Speak() {
	fmt.Println(an.noise)
}

func main() {
	for true {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()
		args := strings.Split(userInput, " ")

		var an Animal

		switch args[0] {
		case "cow":
			an = Animal{food: "grass", locomotion: "walk", noise: "moo"}
		case "bird":
			an = Animal{food: "worms", locomotion: "fly", noise: "peep"}
		case "snake":
			an = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
		default:
			log.Fatalf("Sorry, there are no methods for the '%s'", args[0])
		}

		switch args[1] {
		case "eat":
			an.Eat()
		case "move":
			an.Move()
		case "speak":
			an.Speak()
		default:
			log.Fatalf("Sorry, there are no methods for the '%s'", args[1])
		}
	}
}
