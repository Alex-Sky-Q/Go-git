package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex // chopstick can be used only by one philosopher at a time, hence we need locking
}

type Phil struct {
	leftCS, rightCS *Chopstick
	eatAmount       int // to track how many times each philosopher ate
	num             int // philosopher ID
}

func (p Phil) eat(wg *sync.WaitGroup, mut *sync.Mutex, host *int) {
	for p.eatAmount != 3 {
		start := false
		mut.Lock()
		if *host < 2 {
			*host++
			start = true // we can only start eating if there are less than two other philosophers eating
		}
		mut.Unlock()

		//fmt.Println(*host)
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
		Phils[i] = &Phil{CS[i], CS[(i+1)%5], 0, i + 1}
	}

	host := 0 // to track the requirement that only two philosophers can eat concurrently
	var mut sync.Mutex
	var wg sync.WaitGroup
	wg.Add(5)
	for _, p := range Phils {
		go p.eat(&wg, &mut, &host)
	}
	wg.Wait()
}
