package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Start")

	go func() {
		printPrime("A")
		wg.Done()
	}()

	go func() {
		printPrime("B")
		wg.Done()
	}()

	fmt.Println("Wait finish")
	wg.Wait()

	fmt.Println("Finished")
}

// printPrime displays prime numbers for the first 5000 numbers.
// printPrime is not special. It just requires a little bit more time to complete.
// When we run the program, what we will see are context switches at some point for some particular
// prime number. We cannot predict when the context switch happen. That's why we say the Go's
// scheduler looks and feels very preemptive even though it is a cooperating scheduler.
func printPrime(prefix string) {
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}

		fmt.Printf("%s:%d\n", prefix, outer)
	}

	fmt.Println("Completed", prefix)
}
