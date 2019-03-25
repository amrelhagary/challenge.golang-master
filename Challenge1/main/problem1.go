package main

import (
	"log"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func problem1() {

	log.Printf("problem1: started --------------------------------------------")

	//////////////////////////////////////////////
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//

	// Channel to "send" messages/orders from MAIN routine to subroutines
	cM := make(chan string)

	// Channel to "receive" messages from ALL subroutines (eventually, to ensure all routine are gracefully ended)
	// cS := make(chan string)

	// Initialize the 10 subroutines
	for inx := 0; inx < 10; inx++ {
		wg.Add(1)
		go func(i int) {
			printRandom1(i, cM)
			wg.Done()
		}(inx)
	}

	// Send 100 messages "orders" from Main routine to print 100 random number
	// each message is received/handled by 1 subroutine
	for i := 0; i < 100; i++ {
		cM <- "Print 1 random number"
	}

	// Close the main/order channel, crucial to inform subroutines to END
	close(cM)

	//////////////////////////////////////////////
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	// time.Sleep(5 * time.Second)
	// Each subroutine should send message before it is finished
	// i.e: if there are X subroutines initialized, it is expected to receive X messages
	wg.Wait()
	log.Printf("problem1: finished --------------------------------------------")
}

func printRandom1(slot int, cM chan string) {

	//
	// Do not change 25 into 10!
	//

	for inx := 0; inx < 25; inx++ {
		_, ok := <-cM

		if !ok {
			return
		}

		log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
	}
}
