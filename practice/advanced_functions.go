package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// advancedFunctionsDemo shows advanced function patterns in Go.
func advancedFunctionsDemo() {
	fmt.Println("== function types and passing functions ==")
	applyFunc(5, func(x int) int { return x * 2 })

	fmt.Println("== closures and stateful functions ==")
	adder := makeAdderFunc(10)
	fmt.Println(adder(3), adder(4))

	fmt.Println("== currying style ==")
	multBy2 := curryMultiply(2)
	fmt.Println(multBy2(10))

	fmt.Println("== functional options pattern ==")
	srv := NewServer(WithPort(8080), WithName("demo"))
	fmt.Println("server:", srv.name, srv.port)

	fmt.Println("== method values and expressions ==")
	p := personWorker{name: "x"}
	// method value binds receiver
	mv := p.doWork
	mv("task1")

	fmt.Println("== generic function example (Reduce) ==")
	nums := []int{1, 2, 3, 4}
	sum := Reduce[int](nums, 0, func(a, b int) int { return a + b })
	fmt.Println("sum:", sum)

	fmt.Println("== decorator / retry wrapper ==")
	flaky := func() error {
		if rand.Intn(3) == 0 {
			return nil
		}
		return errors.New("temporary")
	}
	retryable := Retry(flaky, 3, 50*time.Millisecond)
	fmt.Println("retryable result:", retryable())

	fmt.Println("== goroutine + closure capture caution ==")
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		i := i // new variable to capture
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("goroutine i=", i)
		}()
	}
	wg.Wait()

	fmt.Println("== context-aware function ==")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	err := doWorkWithContext(ctx)
	fmt.Println("doWorkWithContext ->", err)
}

// applyFunc accepts a function as a parameter.
func applyFunc(x int, fn func(int) int) {
	fmt.Println("applyFunc result:", fn(x))
}

// makeAdderFunc returns a closure that captures state.
func makeAdderFunc(base int) func(int) int {
	return func(n int) int {
		base += n
		return base
	}
}

// curryMultiply returns a single-arg function bound to multiplier.
func curryMultiply(mult int) func(int) int {
	return func(x int) int { return mult * x }
}

// Functional options pattern
type Server struct {
	name string
	port int
}

type Option func(*Server)

func WithName(n string) Option { return func(s *Server) { s.name = n } }
func WithPort(p int) Option    { return func(s *Server) { s.port = p } }

func NewServer(opts ...Option) *Server {
	s := &Server{name: "default", port: 80}
	for _, o := range opts {
		o(s)
	}
	return s
}

// method value example
type personWorker struct{ name string }

func (p personWorker) doWork(task string) { fmt.Println(p.name, "doing", task) }

// Generic Reduce example (Go 1.18+)
func Reduce[T any](items []T, init T, fn func(T, T) T) T {
	acc := init
	for _, v := range items {
		acc = fn(acc, v)
	}
	return acc
}

// Retry decorator: wraps a func() error with retry logic
func Retry(fn func() error, attempts int, delay time.Duration) func() error {
	return func() error {
		var err error
		for i := 0; i < attempts; i++ {
			err = fn()
			if err == nil {
				return nil
			}
			time.Sleep(delay)
		}
		return err
	}
}

// Context-aware work function
func doWorkWithContext(ctx context.Context) error {
	select {
	case <-time.After(100 * time.Millisecond):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
