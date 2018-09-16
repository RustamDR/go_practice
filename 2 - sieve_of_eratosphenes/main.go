package main

import (
	"fmt"
)

type pipeType struct {
	num  int
	in   chan int
	next chan int
}

type Primes chan int

// Method emulates getting primes using "Sieve Of Eratosphene"
// Returns Primes - chanel with source of primes less or equal limit
func GetPrimes(limit int) Primes {

	allNumbers, primes := make(chan int, 100), make(Primes, 100) // Creating the chanel for all sorted numbers including limit
	go func() {
		check := map[int]struct{}{1: struct{}{}, 3: struct{}{}, 5: struct{}{}, 7: struct{}{}, 9: struct{}{}} // Then filtering only odd numbers
		for i := 3; i <= limit; i++ {
			if _, ok := check[i%10]; ok {
				allNumbers <- i
			}
		}
		close(allNumbers)
	}()

	go addFilter(pipeType{2, allNumbers, primes}) // Starting filtering with 2

	return primes
}

// Launch new pipe with next prime filtering
func addFilter(pipe pipeType) {
	ownPrime, start, primes := pipe.num, true, pipe.next
	primes <- ownPrime

	for num := range pipe.in {
		if num == 0 {
			return
		}

		if num == ownPrime || num%ownPrime == 0 {
			continue
		}

		if start {
			start = false

			pipe.next = make(chan int, 100)             // Creating new filtering chanel by new founded prime
			newPipe := pipeType{num, pipe.next, primes} // Connecting pipes

			go addFilter(newPipe)
		}

		pipe.next <- num
	}

	close(pipe.next)
}

func main() {
	for prime := range GetPrimes(1000000) {
		fmt.Println(prime)
	}
}
