package main

import "sync"

type Chopstick struct {
	sync.Mutex
}

type Phil struct {
	leftCS, rightCS *Chopstick
	eatAmount       int8
}

func main() {
	
}
