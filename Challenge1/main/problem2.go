package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func problem2() {

	log.Printf("problem2: started --------------------------------------------")

	///////////////////////////////////////////////
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//

	c := make(chan int)

	for inx := 0; inx < 10; inx++ {
		wg2.Add(1)
		go func(i int) {
			printRandom2(i, c)
			wg2.Done()
		}(inx)
	}

	// Subroutine (orchestrator) that sends/triggers to channel
	go func() {
		for {
			c <- 0
			time.Sleep(time.Second)
		}
	}()

	///////////////////////////////////////////////
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	//time.Sleep(5 * time.Second)

	wg2.Wait()

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, c chan int) {

	for inx := 0; inx < 10; inx++ {
		_, ok := <-c

		if !ok {
			return
		}

		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

	}
}
