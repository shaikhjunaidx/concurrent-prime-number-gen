package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"
)

func main() {
	var upperLimit int

	flag.IntVar(&upperLimit, "limit", 0, "Upper limit up to which the prime numbers should be generated")

	flag.Parse()

	if upperLimit <= 0 {
		fmt.Println("Please provide a positive integer for the upper limit")
		os.Exit(1)
	}

	fmt.Printf("Generating prime numbers up to %d...\n", upperLimit)

	startTime := time.Now()

	primesChan := make(chan int)
	var wg sync.WaitGroup

	numWorkers := runtime.NumCPU()
	// numWorkers := 4
	fmt.Println("Number of workers:", numWorkers)
	rangeSize := (upperLimit + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize

		if end > upperLimit {
			end = upperLimit
		}

		wg.Add(1)
		go worker(start, end, primesChan, &wg)
	}

	go func() {
		wg.Wait()
		close(primesChan)
	}()

	fmt.Printf("Time taken with goroutines: %v\n", time.Since(startTime))

	var primes []int
	for prime := range primesChan {
		primes = append(primes, prime)
	}

	sort.Ints(primes)

	for _, prime := range primes {
		fmt.Println(prime)
	}


	startTime = time.Now()

	primes = generatePrimesSequential(upperLimit)

	fmt.Printf("Time taken without goroutines: %v\n", time.Since(startTime))

	sort.Ints(primes)
	fmt.Println("Prime numbers:")
	for _, prime := range primes {
		fmt.Println(prime)
	}


}

func generatePrimesSequential(upperLimit int) []int {
	var primes []int
	for i := 2; i <= upperLimit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func worker(start, end int, primesChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primesChan <- i
		}
	}
}
