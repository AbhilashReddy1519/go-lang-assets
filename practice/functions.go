package main

import "fmt"

// functionsDemo demonstrates the main kinds of functions in Go.
// Call this from your main() to see the examples in action.
func functionsDemo() {
	fmt.Println("=== FUNCTION BASICS ===")

	// Basic function call with a simple return value.
	result := add(3, 5)
	fmt.Printf("add(3, 5) = %d\n", result)

	// Multiple return values are common in Go, especially for results + error.
	x, y := swap("first", "second")
	fmt.Printf("swap => x=%q, y=%q\n", x, y)

	// Named return values allow the function to declare result variable names.
	// They are useful when the return values are logically named and the function is short.
	sum, diff := namedReturns(8, 3)
	fmt.Printf("namedReturns => sum=%d, diff=%d\n", sum, diff)

	// Variadic functions accept a variable number of arguments.
	fmt.Printf("sumAll => %d\n", sumAll(1, 2, 3, 4, 5))

	// Function values can be stored in variables and passed around.
	printer := greet
	printer("Abhilash")

	// Higher-order function: pass a function as a parameter.
	once := repeat(3, func(value int) {
		fmt.Printf("repeat value=%d\n", value)
	})
	once(10)

	// Function returning a function (closure).
	adder := makeAdder(10)
	fmt.Printf("makeAdder => %d\n", adder(5))

	// Anonymous function assigned to a variable.
	doWork := func(name string) {
		fmt.Printf("anonymous function running for %s\n", name)
	}
	doWork("Go")

	// Immediately invoked function expression (IIFE).
	func() {
		fmt.Println("immediately invoked anonymous function")
	}()

	// Methods are functions with receivers.
	person := Person{Name: "Nina", Age: 28}
	fmt.Println(person.Greet())
	person.IncAge()
	fmt.Printf("after IncAge => %s is %d\n", person.Name, person.Age)

	// Defer runs a function after the surrounding function returns.
	defer fmt.Println("deferred call executes last")
	fmt.Println("defer example running")
}

// add returns the sum of two integers.
// Use this for simple math or helper functions.
func add(a int, b int) int {
	return a + b
}

// swap returns two values in reversed order.
// Multiple return values are useful for pair operations or tuple-style results.
func swap(a, b string) (string, string) {
	return b, a
}

// namedReturns declares named output variables.
// This style is helpful when a function returns several values and the names aid readability.
func namedReturns(a, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b
	return // returns named values automatically
}

// sumAll accepts any number of integers.
// Use variadic functions when the exact number of inputs is not known.
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// greet is a simple function that can also be used as a function value.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// repeat accepts a function as an argument, demonstrating higher-order functions.
func repeat(times int, fn func(int)) func(int) {
	return func(value int) {
		for i := 0; i < times; i++ {
			fn(value)
		}
	}
}

// makeAdder returns a closure that captures the outer variable base.
// Closures are useful when you need state to persist across calls.
func makeAdder(base int) func(int) int {
	return func(value int) int {
		return base + value
	}
}

// Person is a simple struct to demonstrate methods.
// type Person struct {
// 	Name string
// 	Age  int
// }

// Greet is a method with a value receiver.
// Use value receivers when the method does not need to modify the receiver.
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I am %s", p.Name)
}

// IncAge is a method with a pointer receiver.
// Use pointer receivers when the method modifies the receiver or when the value is large.
func (p *Person) IncAge() {
	p.Age++
}
