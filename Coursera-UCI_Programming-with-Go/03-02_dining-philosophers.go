package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Phil struct {
	leftCS, rightCS *Chopstick
	eatAmount       int
	num             int
}

func (p Phil) eat(wg *sync.WaitGroup, mut *sync.Mutex, host *int) {
	for p.eatAmount != 3 {
		start := false
		mut.Lock()
		if *host < 2 {
			*host++
			start = true
		}
		mut.Unlock()

		fmt.Println(*host)
		if start {
			p.leftCS.Lock()
			p.rightCS.Lock()
			fmt.Printf("Starting to eat %v (cnt: %v) \n", p.num, p.eatAmount+1)
			p.eatAmount++
			fmt.Printf("Finishing eating %v (cnt: %v) \n", p.num, p.eatAmount)
			p.leftCS.Unlock()
			p.rightCS.Unlock()

			mut.Lock()
			*host--
			mut.Unlock()
		}

	}
	wg.Done()
}

func main() {
	CS := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		CS[i] = new(Chopstick)
	}

	Phils := make([]*Phil, 5)
	for i := 0; i < 5; i++ {
		Phils[i] = &Phil{CS[i], CS[(i+1)%5], 0, i}
	}

	host := 0
	var mut sync.Mutex
	var wg sync.WaitGroup
	wg.Add(5)
	for _, p := range Phils {
		go p.eat(&wg, &mut, &host)
	}
	wg.Wait()
}
