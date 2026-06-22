package main

import (
	"fmt"
	"sync"
	"time"
)

/*
GOROUTINES - A Comprehensive Guide

Goroutines are lightweight threads managed by the Go runtime. They are much cheaper
than OS threads and allow you to run thousands or millions concurrently.

Key Differences from Threads:
- Much lighter weight (OS threads use ~1-2MB, goroutines use ~2-8KB)
- Managed by Go runtime, not the OS
- Can have multiple goroutines on a single OS thread
- M:N model - many goroutines can run on few OS threads

Ways to Create and Use Goroutines:
*/

// ============================================================================
// METHOD 1: Simple Goroutine with go keyword
// ============================================================================
func basicGoroutine() {
	fmt.Println("\n--- Method 1: Basic Goroutine ---")

	// go keyword launches a goroutine
	go func() {
		fmt.Println("Running in goroutine 1")
	}()

	// Main function continues immediately (non-blocking)
	fmt.Println("Main continues without waiting")

	// We need to pause main so goroutine has time to execute
	time.Sleep(1 * time.Second)
	fmt.Println("Main completed")
}

// ============================================================================
// METHOD 2: Goroutines with WaitGroup (Synchronization)
// ============================================================================
func goroutineWithWaitGroup() {
	fmt.Println("\n--- Method 2: WaitGroup Synchronization ---")

	var wg sync.WaitGroup

	// Add 3 goroutines to wait for
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		// Launch goroutine
		go func(taskID int) {
			defer wg.Done() // Mark goroutine as done when it exits

			fmt.Printf("Task %d started\n", taskID)
			time.Sleep(time.Duration(taskID) * 500 * time.Millisecond)
			fmt.Printf("Task %d completed\n", taskID)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All tasks completed!")
}

// ============================================================================
// METHOD 3: Goroutines with Channels (Communication)
// ============================================================================
func goroutineWithChannels() {
	fmt.Println("\n--- Method 3: Channel Communication ---")

	// Create a channel to receive results
	results := make(chan string)

	// Launch 3 goroutines that send results
	for i := 1; i <= 3; i++ {
		go func(id int) {
			time.Sleep(time.Duration(id) * 200 * time.Millisecond)
			results <- fmt.Sprintf("Result from goroutine %d", id)
		}(i)
	}

	// Receive results from all goroutines
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}
}

// ============================================================================
// METHOD 4: Worker Pool Pattern (Multiple Goroutines)
// ============================================================================
func workerPoolPattern() {
	fmt.Println("\n--- Method 4: Worker Pool Pattern ---")

	jobs := make(chan int, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// Create 3 worker goroutines
	numWorkers := 3
	wg.Add(numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go func(workerID int) {
			defer wg.Done()

			// Each worker processes jobs from channel
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(200 * time.Millisecond)
				results <- fmt.Sprintf("Job %d completed by worker %d", job, workerID)
			}
		}(w)
	}

	// Send 5 jobs
	go func() {
		for job := 1; job <= 5; job++ {
			jobs <- job
		}
		close(jobs) // Signal workers that no more jobs
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}

// ============================================================================
// METHOD 5: Fan-Out Pattern (One producer, multiple consumers)
// ============================================================================
func fanOutPattern() {
	fmt.Println("\n--- Method 5: Fan-Out Pattern ---")

	// Producer creates tasks
	tasks := make(chan int, 10)

	// Create multiple consumer goroutines
	for i := 1; i <= 3; i++ {
		go func(consumerID int) {
			for task := range tasks {
				fmt.Printf("Consumer %d processing task %d\n", consumerID, task)
			}
		}(i)
	}

	// Send tasks
	for task := 1; task <= 10; task++ {
		tasks <- task
	}
	close(tasks)

	time.Sleep(1 * time.Second)
}

// ============================================================================
// METHOD 6: Fan-In Pattern (Multiple producers, one consumer)
// ============================================================================
func fanInPattern() {
	fmt.Println("\n--- Method 6: Fan-In Pattern ---")

	// Single result channel
	results := make(chan string)

	// Multiple producer goroutines
	for i := 1; i <= 3; i++ {
		go func(producerID int) {
			for j := 1; j <= 2; j++ {
				results <- fmt.Sprintf("Producer %d: Message %d", producerID, j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Single consumer reading from multiple producers
	for i := 0; i < 6; i++ {
		fmt.Println(<-results)
	}
}

// ============================================================================
// METHOD 7: Pipeline Pattern (Data flows through stages)
// ============================================================================
func pipelinePattern() {
	fmt.Println("\n--- Method 7: Pipeline Pattern ---")

	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Stage 2: Double the numbers
	doubled := make(chan int)
	go func() {
		for num := range numbers {
			doubled <- num * 2
		}
		close(doubled)
	}()

	// Stage 3: Square the doubled numbers
	squared := make(chan int)
	go func() {
		for num := range doubled {
			squared <- num * num
		}
		close(squared)
	}()

	// Stage 4: Print results
	for result := range squared {
		fmt.Printf("Final result: %d\n", result)
	}
}

// ============================================================================
// METHOD 8: Goroutine with Context for Cancellation
// ============================================================================
func goroutineWithTimeout() {
	fmt.Println("\n--- Method 8: Goroutine with Timeout ---")

	// Using channels to implement timeout
	done := make(chan bool)
	timeout := time.After(2 * time.Second)

	go func() {
		// Simulating long-running task
		time.Sleep(3 * time.Second)
		done <- true
	}()

	// Wait for either completion or timeout
	select {
	case <-done:
		fmt.Println("Task completed successfully")
	case <-timeout:
		fmt.Println("Task timed out!")
	}
}

// ============================================================================
// METHOD 9: Race Condition Example (Without Mutex - WRONG)
// ============================================================================
func raceConditionExample() {
	fmt.Println("\n--- Method 9: Race Condition (Without Protection) ---")

	counter := 0
	var wg sync.WaitGroup

	// Launch 10 goroutines that increment counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // RACE CONDITION: Multiple goroutines access simultaneously
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Counter (should be 10000, likely less due to race): %d\n", counter)
}

// ============================================================================
// METHOD 10: Multiple Goroutines with Error Handling
// ============================================================================
func goroutineWithErrorHandling() {
	fmt.Println("\n--- Method 10: Goroutines with Error Handling ---")

	type Result struct {
		ID    int
		Value string
		Err   error
	}

	results := make(chan Result, 3)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Simulate work
			time.Sleep(100 * time.Millisecond)

			if id == 2 {
				results <- Result{ID: id, Err: fmt.Errorf("error in goroutine %d", id)}
			} else {
				results <- Result{ID: id, Value: fmt.Sprintf("Success %d", id)}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Err != nil {
			fmt.Printf("Error: %v\n", result.Err)
		} else {
			fmt.Printf("Result: %v\n", result.Value)
		}
	}
}

// Main function demonstrating all goroutine patterns
func demoGoroutines() {
	fmt.Println("========================================")
	fmt.Println("    GOROUTINES - All Patterns Demo")
	fmt.Println("========================================")

	basicGoroutine()
	goroutineWithWaitGroup()
	goroutineWithChannels()
	workerPoolPattern()
	fanOutPattern()
	fanInPattern()
	pipelinePattern()
	goroutineWithTimeout()
	raceConditionExample()
	goroutineWithErrorHandling()
}

// Uncomment the line below to run this demo
// func main() {
// 	demoGoroutines()
// }
