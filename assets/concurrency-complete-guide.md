# Go Concurrency - Complete Guide

A comprehensive guide to mastering concurrency in Go with practical examples and best practices.

---

## Table of Contents

1. [Concurrency vs Parallelism](#concurrency-vs-parallelism)
2. [Goroutines](#goroutines)
3. [Channels](#channels)
4. [Select Statement](#select-statement)
5. [Synchronization Primitives](#synchronization-primitives)
6. [Advanced Patterns](#advanced-patterns)
7. [Best Practices & Common Pitfalls](#best-practices--common-pitfalls)

---

## Concurrency vs Parallelism

### Understanding the Difference

**Concurrency** = Ability to manage multiple tasks by switching between them
**Parallelism** = Actually running multiple tasks at the same time (requires multiple CPU cores)

### Visual Explanation

```
Concurrency (Single Core):
Task 1: [===]    [===]
Task 2:     [===]    [===]
         ↑ Context switches

Parallelism (Multiple Cores):
Core 1: Task 1 [===============]
Core 2: Task 2 [===============]
         ↑ Truly simultaneous
```

### Why Go Excels at Concurrency

1. **Lightweight Goroutines**: Thousands/millions of goroutines on a single machine
2. **Simple Syntax**: `go` keyword is all you need
3. **Built-in Channels**: First-class primitives for communication
4. **Efficient Scheduler**: Go runtime manages goroutines efficiently

### Example: Concurrency in Practice

```go
package main

import (
    "fmt"
    "time"
)

func task(id int) {
    for i := 1; i <= 3; i++ {
        fmt.Printf("Task %d: Step %d\n", id, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // Concurrent execution - tasks interleave
    go task(1)
    go task(2)
    
    time.Sleep(1 * time.Second)
    // Output will show interleaved execution
}
```

---

## Goroutines

### What They Are

Goroutines are lightweight threads managed by the Go runtime. Unlike OS threads:
- **Memory-efficient**: ~2KB per goroutine vs ~1-2MB per OS thread
- **Scheduler-managed**: Runtime decides when to execute
- **Multiplexed**: Many goroutines run on few OS threads

### Why Use Goroutines?

- Write concurrent code that looks sequential
- Easy to create: just prefix with `go` keyword
- Cheap to create thousands/millions
- Natural for I/O-bound operations

### The Scheduler (M:N Model)

Go uses an **M:N scheduler**: M goroutines mapped to N OS threads

```
[Goroutine 1] ─┐
[Goroutine 2] ─┼─→ [OS Thread 1]
[Goroutine 3] ─┘
[Goroutine 4] ─┐
[Goroutine 5] ─┼─→ [OS Thread 2]
[Goroutine 6] ─┘
```

### Goroutine Lifecycle

```
Created → Runnable → Running → Blocked → Runnable → Completed

go myFunc()  ↑         ↑        ↑         ↑
             Start    Scheduling  Wait    Resume
```

### Basic Examples

#### 1. Simple Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func greet(name string) {
    fmt.Printf("Hello, %s!\n", name)
}

func main() {
    // Without goroutine: synchronous
    greet("Alice")        // Waits here
    greet("Bob")          // Then executes
    
    // With goroutines: concurrent
    go greet("Alice")     // Doesn't wait
    go greet("Bob")       // Executes immediately
    
    time.Sleep(100 * time.Millisecond)  // Wait for goroutines
}
```

#### 2. Goroutine with Loop

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int) {
    for i := 0; i < 5; i++ {
        fmt.Printf("Worker %d: Task %d\n", id, i)
        time.Sleep(50 * time.Millisecond)
    }
}

func main() {
    // Create multiple workers
    for i := 1; i <= 3; i++ {
        go worker(i)
    }
    
    time.Sleep(500 * time.Millisecond)
}

// Output (interleaved):
// Worker 1: Task 0
// Worker 2: Task 0
// Worker 3: Task 0
// Worker 1: Task 1
// Worker 2: Task 1
// ...
```

#### 3. Goroutine Lifetime

```go
package main

import (
    "fmt"
    "time"
)

func process() {
    fmt.Println("Starting processing...")
    time.Sleep(2 * time.Second)
    fmt.Println("Processing complete")
}

func main() {
    go process()
    
    fmt.Println("Main continuing...")
    time.Sleep(1 * time.Second)
    
    fmt.Println("Main ending")
    // ⚠️ Program exits! Goroutine is killed mid-processing
    // Output:
    // Main continuing...
    // Starting processing...
    // Main ending
    // (Processing complete is never printed)
}
```

**Key Point**: When `main()` exits, ALL goroutines are terminated, even if unfinished!

---

## Channels

### What Are Channels?

Channels are **typed pipes** for communication between goroutines. They enforce:
- **Type safety**: Can only pass specific types
- **Synchronization**: Sends and receives block until both parties are ready
- **Message passing**: Prefer passing data through channels over sharing memory

### Design Philosophy

> "Share memory by communicating rather than communicate by sharing memory"

### Types of Channels

#### 1. Unbuffered Channels

```go
// Create unbuffered channel
ch := make(chan int)

// Characteristics:
// - Sender blocks until receiver ready
// - Receiver blocks until sender has data
// - Zero capacity
```

**Example: Unbuffered Channel**

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    
    // Goroutine sends data
    go func() {
        fmt.Println("Sending 42...")
        ch <- 42  // Blocks until received
        fmt.Println("Sent!")
    }()
    
    fmt.Println("Main: Waiting to receive...")
    value := <-ch  // Blocks until data available
    fmt.Printf("Received: %v\n", value)
}

// Output:
// Main: Waiting to receive...
// Sending 42...
// Received: 42
// Sent!
```

**Why Unbuffered?** Perfect for strict synchronization and handshakes between goroutines.

#### 2. Buffered Channels

```go
// Create buffered channel with capacity of 3
ch := make(chan int, 3)

// Characteristics:
// - Sender only blocks if buffer full
// - Receiver only blocks if buffer empty
// - Can accept N values before blocking
```

**Example: Buffered Channel**

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)  // Buffer size: 2
    
    // These don't block because buffer has space
    ch <- 1
    ch <- 2
    fmt.Println("Both values sent")
    
    // This would block (buffer full)
    // ch <- 3  // Would wait here
    
    // Receive values
    fmt.Println(<-ch)  // 1
    fmt.Println(<-ch)  // 2
}

// Output:
// Both values sent
// 1
// 2
```

**Why Buffered?** Decouple sender and receiver, batch operations, rate limiting.

#### 3. Channel Operations

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    
    go func() {
        ch <- 10
        close(ch)  // Signal no more data
    }()
    
    // Receive with status check
    value, ok := <-ch
    fmt.Printf("Value: %v, Channel open: %v\n", value, ok)
    // Value: 10, Channel open: true
    
    value, ok = <-ch
    fmt.Printf("Value: %v, Channel open: %v\n", value, ok)
    // Value: 0, Channel open: false
}
```

#### 4. Closing Channels

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    
    // Send values
    ch <- 1
    ch <- 2
    ch <- 3
    
    // Close the channel
    close(ch)
    
    // Range over closed channel (drains remaining values)
    for value := range ch {
        fmt.Println(value)  // 1, 2, 3
    }
    
    fmt.Println("Channel exhausted")
}

// Rules for closing:
// 1. Only sender should close (panic if receiver closes)
// 2. Receiving from closed channel returns zero value + false
// 3. Can range over closed channel
// 4. Sending to closed channel panics
```

### Real-World Channel Pattern: Producer-Consumer

```go
package main

import (
    "fmt"
    "time"
)

// Producer sends data
func producer(ch chan string) {
    for i := 1; i <= 5; i++ {
        data := fmt.Sprintf("Data-%d", i)
        fmt.Printf("Producing: %s\n", data)
        ch <- data
        time.Sleep(200 * time.Millisecond)
    }
    close(ch)  // Signal completion
}

// Consumer receives data
func consumer(ch chan string) {
    for data := range ch {  // Auto-exits when closed
        fmt.Printf("Consuming: %s\n", data)
        time.Sleep(300 * time.Millisecond)
    }
    fmt.Println("Consumer finished")
}

func main() {
    ch := make(chan string)
    
    go producer(ch)
    go consumer(ch)
    
    time.Sleep(3 * time.Second)
}

// Output shows producer ahead, then consumer catches up
```

---

## Select Statement

### What Is Select?

`select` is like `switch` for channels - it waits on multiple channel operations.

### When to Use

- Listen to multiple channels
- Implement timeouts
- Non-blocking channel operations
- Choose first ready operation

### Basic Syntax

```go
select {
case value := <-ch1:
    // Handle ch1
    
case value := <-ch2:
    // Handle ch2
    
case ch3 <- value:
    // Send to ch3
    
default:
    // If no case ready (optional)
}
```

### Example 1: Multiple Channels

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "Result from channel 1"
    }()
    
    go func() {
        time.Sleep(50 * time.Millisecond)
        ch2 <- "Result from channel 2"
    }()
    
    // Wait for first result
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Got from ch1:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Got from ch2:", msg2)
        }
    }
}

// Output:
// Got from ch2: Result from channel 2
// Got from ch1: Result from channel 1
```

### Example 2: Timeout Pattern

```go
package main

import (
    "fmt"
    "time"
)

func slowOperation(result chan string) {
    time.Sleep(3 * time.Second)
    result <- "Done!"
}

func main() {
    result := make(chan string)
    go slowOperation(result)
    
    // Implement timeout
    select {
    case msg := <-result:
        fmt.Println("Success:", msg)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout! Operation took too long")
    }
}

// Output:
// Timeout! Operation took too long
```

### Example 3: Non-Blocking Operations

```go
package main

import "fmt"

func main() {
    ch := make(chan int)
    
    select {
    case value := <-ch:
        fmt.Println("Received:", value)
    default:
        fmt.Println("No data available (non-blocking)")
    }
    
    // Send non-blocking
    select {
    case ch <- 42:
        fmt.Println("Sent")
    default:
        fmt.Println("Channel not ready (non-blocking)")
    }
}

// Output:
// No data available (non-blocking)
// Channel not ready (non-blocking)
```

### Example 4: Select in Loop (Multiplexing)

```go
package main

import (
    "fmt"
    "time"
)

func task(id int, ch chan string) {
    for i := 0; i < 3; i++ {
        time.Sleep(time.Duration(id*100) * time.Millisecond)
        ch <- fmt.Sprintf("Task %d: Message %d", id, i)
    }
    close(ch)
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go task(1, ch1)
    go task(2, ch2)
    
    count := 0
    for count < 2 {
        select {
        case msg, ok := <-ch1:
            if !ok {
                count++
            } else {
                fmt.Println(msg)
            }
        case msg, ok := <-ch2:
            if !ok {
                count++
            } else {
                fmt.Println(msg)
            }
        }
    }
}
```

---

## Synchronization Primitives

### WaitGroups

**Purpose**: Coordinate goroutines, wait for group to complete

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Mark this goroutine as done
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Duration(id*100) * time.Millisecond)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)  // Increment counter
        go worker(i, &wg)
    }
    
    wg.Wait()  // Block until counter reaches 0
    fmt.Println("All workers completed")
}

// Output:
// Worker 1 starting
// Worker 2 starting
// Worker 3 starting
// Worker 1 done
// Worker 2 done
// Worker 3 done
// All workers completed
```

**When to Use**: When you need to wait for multiple goroutines, but don't need communication between them.

### Mutex (Mutual Exclusion)

**Purpose**: Protect shared data from concurrent access

#### sync.Mutex

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    return c.value
}

func main() {
    counter := &Counter{}
    var wg sync.WaitGroup
    
    // 100 goroutines incrementing
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final count: %d\n", counter.Value())
    // Without mutex: unpredictable (might be < 100)
    // With mutex: guaranteed 100
}
```

**Why Mutex?** Prevents race conditions when multiple goroutines access shared data.

#### sync.RWMutex (Reader-Writer Mutex)

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Cache struct {
    mu    sync.RWMutex
    data  map[string]string
}

func (c *Cache) Set(key, value string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    c.data[key] = value
    fmt.Printf("Wrote: %s = %s\n", key, value)
}

func (c *Cache) Get(key string) string {
    c.mu.RLock()  // Read lock (shared)
    defer c.mu.RUnlock()
    
    value := c.data[key]
    fmt.Printf("Read: %s = %s\n", key, value)
    return value
}

func main() {
    cache := &Cache{data: make(map[string]string)}
    
    // Multiple readers don't block each other
    for i := 0; i < 3; i++ {
        go func() {
            for j := 0; j < 2; j++ {
                cache.Get("key")
                time.Sleep(50 * time.Millisecond)
            }
        }()
    }
    
    // Writers block both readers and other writers
    go func() {
        time.Sleep(75 * time.Millisecond)
        cache.Set("key", "value")
    }()
    
    time.Sleep(500 * time.Millisecond)
}
```

**Difference**:
- `Mutex`: Exclusive access (one writer OR nothing)
- `RWMutex`: Multiple readers OR one writer (better for read-heavy workloads)

### Atomic Operations

**Purpose**: Thread-safe operations without explicit locks

```go
package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var count int64
    var wg sync.WaitGroup
    
    // 1000 goroutines increment atomically
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&count, 1)
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final count: %d\n", count)
    // Guaranteed: 1000 (no race condition)
}
```

**When to Use**: Simple counter/flag operations (faster than mutex).

---

## Advanced Patterns

### Worker Pool Pattern

**Problem**: Creating unlimited goroutines is wasteful

**Solution**: Fixed pool of workers processing jobs

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID    int
    Data  string
}

type Result struct {
    Job    Job
    Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        fmt.Printf("Worker %d processing Job %d\n", id, job.ID)
        time.Sleep(500 * time.Millisecond)
        
        results <- Result{
            Job:    job,
            Output: fmt.Sprintf("Result for %s", job.Data),
        }
    }
}

func main() {
    numWorkers := 3
    numJobs := 10
    
    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)
    
    // Start workers
    for w := 1; w <= numWorkers; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{ID: j, Data: fmt.Sprintf("Job-%d", j)}
    }
    close(jobs)
    
    // Collect results
    for r := 0; r < numJobs; r++ {
        result := <-results
        fmt.Printf("Got result: %s\n", result.Output)
    }
}

// Benefits:
// - Reusable workers (no overhead creating goroutines per job)
// - Backpressure (jobs queue if workers busy)
// - Easy to scale (adjust numWorkers)
```

### Fan-Out / Fan-In Pattern

**Fan-Out**: One input → multiple outputs
**Fan-In**: Multiple inputs → one output

```go
package main

import (
    "fmt"
    "sync"
)

// Fan-Out: Split work among workers
func fanOut(input <-chan int, numWorkers int) []<-chan int {
    channels := make([]<-chan int, numWorkers)
    
    for i := 0; i < numWorkers; i++ {
        ch := make(chan int)
        channels[i] = ch
        
        go func(out chan int) {
            defer close(out)
            for val := range input {
                out <- val * val  // Square each number
            }
        }(ch)
    }
    
    return channels
}

// Fan-In: Merge multiple channels
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for val := range c {
                out <- val
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

func main() {
    // Create input
    input := make(chan int)
    go func() {
        for i := 1; i <= 5; i++ {
            input <- i
        }
        close(input)
    }()
    
    // Fan-out to 2 workers
    outputs := fanOut(input, 2)
    
    // Fan-in results
    results := fanIn(outputs...)
    
    // Consume results
    for result := range results {
        fmt.Println(result)  // 1, 4, 9, 16, 25 (in some order)
    }
}
```

### Pipeline Pattern

**Concept**: Chain of processing stages connected by channels

```go
package main

import (
    "fmt"
)

// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Stage 3: Filter even numbers
func filterEven(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if n%2 == 0 {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

func main() {
    // Build pipeline: generate → square → filter
    input := generate(1, 2, 3, 4, 5)
    squared := square(input)
    filtered := filterEven(squared)
    
    // Consume results
    for result := range filtered {
        fmt.Println(result)  // 4, 16
    }
}

// Pipeline Benefits:
// - Modular: Each stage independent
// - Composable: Easy to add/remove stages
// - Parallel: Stages run concurrently
```

### Context Pattern

**Purpose**: Cancel operations, set deadlines, pass values

```go
package main

import (
    "context"
    "fmt"
    "time"
)

// Cancellation example
func cancellationExample() {
    ctx, cancel := context.WithCancel(context.Background())
    
    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("Worker cancelled")
                return
            default:
                fmt.Println("Working...")
                time.Sleep(200 * time.Millisecond)
            }
        }
    }()
    
    time.Sleep(700 * time.Millisecond)
    cancel()  // Signal cancellation
    time.Sleep(500 * time.Millisecond)
}

// Deadline example
func deadlineExample() {
    ctx, cancel := context.WithDeadline(
        context.Background(),
        time.Now().Add(1 * time.Second),
    )
    defer cancel()
    
    select {
    case <-time.After(2 * time.Second):
        fmt.Println("Task completed")
    case <-ctx.Done():
        fmt.Printf("Deadline exceeded: %v\n", ctx.Err())
    }
}

// Timeout example
func timeoutExample() {
    ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
    defer cancel()
    
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("Operation completed")
    case <-ctx.Done():
        fmt.Printf("Timeout: %v\n", ctx.Err())
    }
}

// Value passing
func valuePassingExample() {
    ctx := context.WithValue(context.Background(), "user_id", 123)
    
    userID := ctx.Value("user_id")
    fmt.Printf("User ID: %v\n", userID)
}

func main() {
    fmt.Println("=== Cancellation ===")
    cancellationExample()
    
    fmt.Println("\n=== Deadline ===")
    deadlineExample()
    
    fmt.Println("\n=== Timeout ===")
    timeoutExample()
    
    fmt.Println("\n=== Value Passing ===")
    valuePassingExample()
}
```

---

## Best Practices & Common Pitfalls

### Race Conditions

**What**: Multiple goroutines accessing shared data unsafely

```go
// ❌ BAD: Race condition
var count int
go func() { count++ }()
go func() { count++ }()
// Result: unpredictable (might be 1 or 2)

// ✅ GOOD: Thread-safe with mutex
var mu sync.Mutex
var count int
go func() {
    mu.Lock()
    count++
    mu.Unlock()
}()
go func() {
    mu.Lock()
    count++
    mu.Unlock()
}()
// Result: guaranteed 2

// ✅ GOOD: Atomic operation
var count int64
go func() { atomic.AddInt64(&count, 1) }()
go func() { atomic.AddInt64(&count, 1) }()
// Result: guaranteed 2
```

**Detection**: Run with `-race` flag
```bash
go run -race main.go
go test -race ./...
```

### Deadlocks

**What**: Goroutines blocked waiting for each other

```go
// ❌ BAD: Deadlock
func main() {
    ch := make(chan int)
    ch <- 1  // Blocked! No receiver
    // <-ch
}

// ✅ GOOD: Receiver ready
func main() {
    ch := make(chan int)
    go func() {
        ch <- 1
    }()
    <-ch
}

// ❌ BAD: Mutual blocking
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        <-ch1
        ch2 <- 1
    }()
    
    go func() {
        <-ch2
        ch1 <- 1
    }()
    // Neither can proceed - deadlock!
}
```

### Goroutine Leaks

**What**: Goroutines that never terminate

```go
// ❌ BAD: Goroutine leak
func fetchData(url string) <-chan string {
    ch := make(chan string)
    go func() {
        // If caller doesn't read from ch, goroutine blocked forever
        ch <- "data from " + url
    }()
    return ch
}

// ✅ GOOD: Use context for cancellation
func fetchData(ctx context.Context, url string) <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch)
        select {
        case ch <- "data from " + url:
        case <-ctx.Done():
            return
        }
    }()
    return ch
}

// ✅ GOOD: Use buffered channel
func fetchData(url string) <-chan string {
    ch := make(chan string, 1)  // Buffered
    go func() {
        ch <- "data from " + url
    }()
    return ch
}
```

### Channel Best Practices

```go
// 1. Close channels only from sender side
func (s *Sender) send(ch chan int) {
    ch <- 42
    close(ch)  // ✅ OK
}

func (r *Receiver) receive(ch chan int) {
    x := <-ch
    close(ch)  // ❌ PANIC! (if multiple senders)
}

// 2. Prefer range over manual checking
// ❌ Manual
for {
    val, ok := <-ch
    if !ok {
        break
    }
    // process val
}

// ✅ Range
for val := range ch {
    // process val
}

// 3. Use directional channels in function signatures
func send(ch chan<- int, value int) {  // Send-only
    ch <- value
}

func receive(ch <-chan int) int {  // Receive-only
    return <-ch
}

// 4. Unbuffered for synchronization, buffered for decoupling
// Synchronization
sync := make(chan struct{})  // Unbuffered

// Decoupling/rate-limiting
work := make(chan Job, 100)  // Buffered
```

### Concurrency Best Practices

```go
// 1. Keep goroutines alive with WaitGroup
var wg sync.WaitGroup
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // work
    }()
}
wg.Wait()

// 2. Use context for graceful shutdown
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

select {
case <-ctx.Done():
    fmt.Println("Timeout")
case result := <-workCh:
    fmt.Println(result)
}

// 3. Avoid globals, pass through channels/params
// ❌ BAD
var sharedData = make(map[string]int)

// ✅ GOOD
func worker(data map[string]int, ch <-chan Job) {
    // work with passed data
}

// 4. Monitor goroutines
func main() {
    fmt.Printf("Goroutines before: %d\n", runtime.NumGoroutine())
    // ... code ...
    fmt.Printf("Goroutines after: %d\n", runtime.NumGoroutine())
}
```

---

## Phase 2: Intermediate Go - Related Topics

### Goroutines Summary
- Lightweight concurrent units managed by Go runtime
- Use `go` keyword to launch
- Always ensure clean termination (WaitGroup, context)
- Monitor with `runtime.NumGoroutine()`

### Channels Summary
- Use for communication between goroutines
- Unbuffered: synchronization | Buffered: decoupling
- Only sender closes
- Range over to drain safely

### Context Summary
- Propagate cancellation signals
- Set deadlines and timeouts
- Pass request-scoped values
- Always use `defer cancel()` with timeout contexts

### Concurrency Summary
- Mutex/RWMutex for shared state
- Atomic for simple values
- WaitGroup for coordination
- Design to prevent race conditions, deadlocks, leaks

### Packages
Covered in separate guide, but concurrency-relevant:
- `sync`: Mutex, RWMutex, WaitGroup, Once
- `sync/atomic`: Atomic counters
- `context`: Cancellation and timeouts
- `time`: Delays, timeouts
- `runtime`: Goroutine monitoring

### Modules
- `go mod init`: Create module
- `go mod tidy`: Clean dependencies
- `go mod download`: Get dependencies
- Version control with `go.mod` and `go.sum`

---

## Real-World Example: Concurrent HTTP Fetcher

```go
package main

import (
    "context"
    "fmt"
    "net/http"
    "sync"
    "time"
)

type Result struct {
    URL    string
    Status int
    Err    error
}

func fetchURL(ctx context.Context, url string, results chan<- Result) {
    client := &http.Client{Timeout: 5 * time.Second}
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        results <- Result{URL: url, Err: err}
        return
    }
    
    resp, err := client.Do(req)
    if err != nil {
        results <- Result{URL: url, Err: err}
        return
    }
    defer resp.Body.Close()
    
    results <- Result{URL: url, Status: resp.StatusCode}
}

func main() {
    urls := []string{
        "https://golang.org",
        "https://github.com",
        "https://google.com",
    }
    
    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    // Buffer to hold results
    results := make(chan Result, len(urls))
    var wg sync.WaitGroup
    
    // Launch fetchers
    for _, url := range urls {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            fetchURL(ctx, u, results)
        }(url)
    }
    
    // Wait for all to complete
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect and display results
    for result := range results {
        if result.Err != nil {
            fmt.Printf("%s: Error - %v\n", result.URL, result.Err)
        } else {
            fmt.Printf("%s: Status %d\n", result.URL, result.Status)
        }
    }
}
```

---

## Key Takeaways

1. **Concurrency != Parallelism**: Concurrency is about managing multiple tasks; parallelism is about running them simultaneously
2. **Goroutines are cheap**: Use them liberally for I/O-bound work
3. **Channels for communication**: Prefer passing data through channels over sharing memory
4. **Context for cancellation**: Use for timeouts and graceful shutdown
5. **Synchronize properly**: Use Mutex, WaitGroup, Atomic appropriately
6. **Test with `-race`**: Always check for race conditions
7. **Monitor goroutines**: Ensure no leaks
8. **Pattern selection matters**: Choose Worker Pool, Fan-In/Fan-Out, or Pipeline based on your use case

---

## Resources for Deeper Learning

- `go run -race main.go`: Detect race conditions
- `go test -race ./...`: Test with race detector
- `runtime.NumGoroutine()`: Monitor active goroutines
- `pprof`: Profile goroutine usage
- [Concurrency in Go (Kathleen Reilly)](https://www.oreilly.com/library/view/concurrency-in-go/9781491941874/)

---

**Happy concurrent coding! 🚀**
