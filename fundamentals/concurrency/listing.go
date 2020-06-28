package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {


	fmt.Println("Create goroutines")
	
	// wg will wait for 2 goroutines to complete
	wg.Add(2)

	go printPrime("A")
	go printPrime("B")

	// Block main until goroutines are done
	wg.Wait()

	fmt.Println("Terminating program")

}

func printPrime(prefix string){
	// Make sure wg counter is decremented when this function returns
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer % inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed:", prefix)
}