package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Basic channel example
	fmt.Println("Example 1: Basic channel")
	ch := make(chan string) // Create an unbuffered channel

	go func() {
		ch <- "Hello, channels!" // Send data to channel
		// This blocks until someone receives
	}()

	msg := <-ch // Receive data from channel
	fmt.Println(msg)

	// Buffered channel example
	fmt.Println("\nExample 2: Buffered channel")
	bufferedCh := make(chan int, 3) // Create a buffered channel with capacity 3

	// Send values without blocking until buffer is full
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	// bufferedCh <- 4  // This would block as buffer is full

	// Receive values
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)
	fmt.Println(<-bufferedCh)

	// Range over channel example
	fmt.Println("\nExample 3: Range over channel")
	numCh := make(chan int)

	// Send values and close channel
	go func() {
		for i := 1; i <= 5; i++ {
			numCh <- i
		}
		close(numCh) // Important: close the channel when done sending
	}()

	// Range receives values until channel is closed
	for num := range numCh {
		fmt.Println("Received:", num)
	}

	// Select statement example
	fmt.Println("\nExample 4: Select statement")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	// Select picks the first channel that's ready
	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received", msg)
		case msg := <-ch2:
			fmt.Println("Received", msg)
		}
	}

	// Worker pool example
	fmt.Println("\nExample 5: Worker pool")
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 worker goroutines
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs) // Signal no more jobs

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
		results <- j * 2 // Send result back
	}
}
