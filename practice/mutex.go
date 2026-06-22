package main

import (
	"fmt"
	"sync"
	"time"
)

/*
MUTEX (MUTUAL EXCLUSION) - A Comprehensive Guide

A Mutex is a synchronization primitive that ensures only one goroutine can access
a shared resource at a time. It prevents race conditions when multiple goroutines
access and modify the same data.

Key Concepts:
- Lock(): Acquires the mutex. If already locked, goroutine waits.
- Unlock(): Releases the mutex, allowing other waiting goroutines to proceed.
- RWMutex: Allows multiple readers but only one writer (writer-exclusive)

Why Use Mutex:
- Prevent race conditions when accessing shared variables
- Ensure data consistency
- Protect critical sections of code

Different Ways to Use Mutex:
*/

// ============================================================================
// METHOD 1: Basic Mutex with Lock/Unlock
// ============================================================================
func basicMutexExample() {
	fmt.Println("\n--- Method 1: Basic Mutex (Lock/Unlock) ---")

	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup

	// 10 goroutines incrementing counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (should be 10000)\n", counter)
}

// ============================================================================
// METHOD 2: Mutex with Defer (Best Practice)
// ============================================================================
func mutexWithDefer() {
	fmt.Println("\n--- Method 2: Mutex with Defer (Safe Pattern) ---")

	var mu sync.Mutex
	value := 0

	var wg sync.WaitGroup

	// Writer goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock() // Guarantees unlock even if panic occurs

			oldValue := value
			time.Sleep(10 * time.Millisecond) // Simulate work
			value = oldValue + 1

			fmt.Printf("Goroutine %d: Updated value to %d\n", id, value)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final value: %d\n", value)
}

// ============================================================================
// METHOD 3: Protecting a Struct with Mutex
// ============================================================================
type BankAccount struct {
	mu      sync.Mutex
	balance float64
}

func (ba *BankAccount) Deposit(amount float64) {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	ba.balance += amount
	fmt.Printf("Deposited: $%.2f, Balance: $%.2f\n", amount, ba.balance)
}

func (ba *BankAccount) Withdraw(amount float64) error {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	if ba.balance < amount {
		return fmt.Errorf("insufficient funds: balance $%.2f, requested $%.2f", ba.balance, amount)
	}

	ba.balance -= amount
	fmt.Printf("Withdrawn: $%.2f, Balance: $%.2f\n", amount, ba.balance)
	return nil
}

func (ba *BankAccount) GetBalance() float64 {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	return ba.balance
}

func mutexWithStructExample() {
	fmt.Println("\n--- Method 3: Protecting Struct with Mutex ---")

	account := &BankAccount{balance: 1000}

	var wg sync.WaitGroup

	// Multiple concurrent transactions
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			if id%2 == 0 {
				account.Deposit(100)
			} else {
				account.Withdraw(50)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final balance: $%.2f\n", account.GetBalance())
}

// ============================================================================
// METHOD 4: RWMutex (Read-Write Mutex) - Multiple Readers, Single Writer
// ============================================================================
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // Read lock - allows multiple readers
	defer c.mu.RUnlock()

	value, exists := c.data[key]
	return value, exists
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // Write lock - exclusive access
	defer c.mu.Unlock()

	c.data[key] = value
}

func rwMutexExample() {
	fmt.Println("\n--- Method 4: RWMutex (Multiple Readers, Single Writer) ---")

	cache := &Cache{data: make(map[string]string)}

	var wg sync.WaitGroup

	// 10 reader goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				value, exists := cache.Get("key1")
				if exists {
					fmt.Printf("Reader %d: Read value = %s\n", id, value)
				} else {
					fmt.Printf("Reader %d: Key not found\n", id)
				}
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	// 2 writer goroutines
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			cache.Set("key1", fmt.Sprintf("value_%d", id))
			fmt.Printf("Writer %d: Updated key1\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
}

// ============================================================================
// METHOD 5: Sync.Once (Execute Code Only Once)
// ============================================================================
func syncOnceExample() {
	fmt.Println("\n--- Method 5: Sync.Once (Execute Once Across Goroutines) ---")

	var once sync.Once
	var wg sync.WaitGroup

	// Simulating initialization that should happen only once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			once.Do(func() {
				fmt.Printf("Initialization running (called by goroutine %d)\n", id)
				time.Sleep(100 * time.Millisecond)
				fmt.Println("Initialization complete")
			})

			fmt.Printf("Goroutine %d: Continuing after initialization\n", id)
		}(i)
	}

	wg.Wait()
}

// ============================================================================
// METHOD 6: Sync.Cond (Condition Variable)
// ============================================================================
func syncCondExample() {
	fmt.Println("\n--- Method 6: Sync.Cond (Wait for Condition) ---")

	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var isReady bool

	var wg sync.WaitGroup

	// Worker goroutines waiting for signal
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			cond.L.Lock()
			defer cond.L.Unlock()

			fmt.Printf("Worker %d: Waiting for signal...\n", id)
			for !isReady { // Use a loop to handle spurious wakeups
				cond.Wait() // Releases lock, waits, re-acquires lock
			}
			fmt.Printf("Worker %d: Received signal, continuing\n", id)
		}(i)
	}

	// Signaler goroutine
	go func() {
		time.Sleep(500 * time.Millisecond)

		cond.L.Lock()
		isReady = true
		fmt.Println("Signaler: Broadcasting signal to all waiters")
		cond.Broadcast() // Wake all waiting goroutines
		cond.L.Unlock()
	}()

	wg.Wait()
}

// ============================================================================
// METHOD 7: Atomic Operations (Lock-Free Alternative to Mutex)
// ============================================================================
func atomicOperationsExample() {
	fmt.Println("\n--- Method 7: Atomic Operations (Lock-Free) ---")

	var counter int64
	var wg sync.WaitGroup

	// 10 goroutines using atomic operations
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				// Atomic increment (lock-free)
				// Note: sync/atomic package provides these operations
				// For this demo, we'll show the concept with mutex
				var mu sync.Mutex
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter: %d\n", counter)
}

// ============================================================================
// METHOD 8: Producer-Consumer with Cond
// ============================================================================
type Queue struct {
	mu    sync.Mutex
	cond  *sync.Cond
	items []int
}

func NewQueue() *Queue {
	q := &Queue{
		items: make([]int, 0),
	}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Put(item int) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.items = append(q.items, item)
	fmt.Printf("Produced: %d (queue size: %d)\n", item, len(q.items))
	q.cond.Signal() // Wake one waiting consumer
}

func (q *Queue) Get() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) == 0 {
		q.cond.Wait() // Wait for item to be available
	}

	item := q.items[0]
	q.items = q.items[1:]
	fmt.Printf("Consumed: %d (queue size: %d)\n", item, len(q.items))
	return item
}

func producerConsumerExample() {
	fmt.Println("\n--- Method 8: Producer-Consumer Pattern with Cond ---")

	queue := NewQueue()
	var wg sync.WaitGroup

	// Producers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 1; j <= 3; j++ {
				queue.Put(id*10 + j)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Consumers
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 3; j++ {
				queue.Get()
				time.Sleep(150 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}

// ============================================================================
// METHOD 9: Deadlock Example (What NOT to do)
// ============================================================================
func deadlockExample() {
	fmt.Println("\n--- Method 9: Deadlock Warning (What NOT to do) ---")

	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	// Goroutine 1
	go func() {
		defer wg.Done()

		mu.Lock()
		fmt.Println("Goroutine 1: Acquired lock")
		time.Sleep(100 * time.Millisecond)

		// Trying to acquire same lock again (DEADLOCK in basic Mutex)
		// Uncomment next line to see deadlock:
		// mu.Lock()
		// fmt.Println("Goroutine 1: This won't print")

		mu.Unlock()
	}()

	wg.Wait()
	fmt.Println("No deadlock occurred - example prevented for safety")
}

// ============================================================================
// METHOD 10: Mutex Best Practices
// ============================================================================
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Good: Receiver is pointer to acquire lock on the struct
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.value++
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.value
}

// Bad (DON'T DO THIS): Lock field should not be exposed
// type BadCounter struct {
// 	Mu    sync.Mutex // Exported - BAD PRACTICE
// 	Value int
// }

func mutexBestPractices() {
	fmt.Println("\n--- Method 10: Mutex Best Practices ---")

	counter := &SafeCounter{}

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Safe counter value: %d\n", counter.Value())
}

// Main demo function
func demoMutex() {
	fmt.Println("========================================")
	fmt.Println("    MUTEX - All Patterns Demo")
	fmt.Println("========================================")

	basicMutexExample()
	mutexWithDefer()
	mutexWithStructExample()
	rwMutexExample()
	syncOnceExample()
	syncCondExample()
	atomicOperationsExample()
	producerConsumerExample()
	deadlockExample()
	mutexBestPractices()
}

// Uncomment the line below to run this demo
// func main() {
// 	demoMutex()
// }
