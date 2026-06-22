package main

import (
	"fmt"
)

/*
GENERICS IN GO - A Comprehensive Guide (Go 1.18+)

Generics allow you to write functions and types that work with any type while
maintaining compile-time type safety. Before Go 1.18, developers had to use
interface{} (empty interface) which lost type information.

Key Concepts:
- Type Parameters: Similar to type variables in other languages
- Constraints: Define what types are allowed (similar to trait bounds in Rust)
- Comparable: Built-in constraint for types that support == and !=
- Ordered: Custom constraint for types that support <, >, <=, >=

Syntax: func FuncName[T Type](param T) T

Ways to Use Generics:
*/

// ============================================================================
// METHOD 1: Basic Generic Function with Single Type Parameter
// ============================================================================

// Generic function that works with any type
func Print[T any](value T) {
	fmt.Printf("Value: %v (Type: %T)\n", value, value)
}

// Generic function that swaps two values
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func basicGenericsExample() {
	fmt.Println("\n--- Method 1: Basic Generic Functions ---")

	Print(42)
	Print("hello")
	Print(3.14)
	Print([]int{1, 2, 3})

	fmt.Println("\nSwapping values:")
	a, b := Swap(10, 20)
	fmt.Printf("After swap: a=%d, b=%d\n", a, b)

	x, y := Swap("hello", "world")
	fmt.Printf("After swap: x=%s, y=%s\n", x, y)
}

// ============================================================================
// METHOD 2: Generic Function with Constraints
// ============================================================================

// Constraint: Only numeric types
type Number interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

// Generic function with Number constraint
func Add[T Number](a, b T) T {
	return a + b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func constrainedGenericsExample() {
	fmt.Println("\n--- Method 2: Generic Functions with Constraints ---")

	fmt.Printf("Add(10, 20) = %d\n", Add(10, 20))
	fmt.Printf("Add(10.5, 20.3) = %f\n", Add(10.5, 20.3))
	fmt.Printf("Max(100, 50) = %d\n", Max(100, 50))
	fmt.Printf("Max(3.14, 2.71) = %f\n", Max(3.14, 2.71))

	// Compiler error if uncommented: string is not in Number constraint
	// fmt.Println(Add("hello", "world"))
}

// ============================================================================
// METHOD 3: Generic Slice Operations
// ============================================================================

// Find checks if a value exists in a slice
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Map applies a function to each element
func MapSlice[T any, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = transform(v)
	}
	return result
}

// Filter keeps only elements that match the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Reduce combines all elements into a single value
func Reduce[T any, U any](slice []T, initial U, reduce func(U, T) U) U {
	result := initial
	for _, item := range slice {
		result = reduce(result, item)
	}
	return result
}

func genericSliceOperationsExample() {
	fmt.Println("\n--- Method 3: Generic Slice Operations ---")

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("Contains(numbers, 3): %v\n", Contains(numbers, 3))
	fmt.Printf("Contains(numbers, 10): %v\n", Contains(numbers, 10))

	// Map: Convert int to string
	strings := MapSlice(numbers, func(n int) string {
		return fmt.Sprintf("num_%d", n)
	})
	fmt.Printf("Mapped to strings: %v\n", strings)

	// Filter: Get even numbers
	evens := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Even numbers: %v\n", evens)

	// Reduce: Sum all numbers
	sum := Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Printf("Sum: %d\n", sum)
}

// ============================================================================
// METHOD 4: Generic Data Structures (Stack Example)
// ============================================================================

// Stack is a generic LIFO data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func genericStackExample() {
	fmt.Println("\n--- Method 4: Generic Data Structure (Stack) ---")

	// Stack of integers
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Printf("Stack size: %d\n", intStack.Size())

	for !intStack.IsEmpty() {
		item, _ := intStack.Pop()
		fmt.Printf("Popped: %d\n", item)
	}

	// Stack of strings
	stringStack := Stack[string]{}
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")

	for !stringStack.IsEmpty() {
		item, _ := stringStack.Pop()
		fmt.Printf("Popped: %s\n", item)
	}
}

// ============================================================================
// METHOD 5: Generic Queue Data Structure
// ============================================================================

// Queue is a generic FIFO data structure
type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Front() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func genericQueueExample() {
	fmt.Println("\n--- Method 5: Generic Data Structure (Queue) ---")

	queue := Queue[string]{}
	queue.Enqueue("first")
	queue.Enqueue("second")
	queue.Enqueue("third")

	for !queue.IsEmpty() {
		item, _ := queue.Dequeue()
		fmt.Printf("Dequeued: %s\n", item)
	}
}

// ============================================================================
// METHOD 6: Generic Pair/Tuple Structure
// ============================================================================

// Pair holds two values of potentially different types
type Pair[T, U any] struct {
	First  T
	Second U
}

func (p Pair[T, U]) Swap() Pair[U, T] {
	return Pair[U, T]{
		First:  p.Second,
		Second: p.First,
	}
}

func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

func genericPairExample() {
	fmt.Println("\n--- Method 6: Generic Pair/Tuple Structure ---")

	// Pair of int and string
	p1 := Pair[int, string]{First: 42, Second: "answer"}
	fmt.Printf("Original: %s\n", p1)
	fmt.Printf("Swapped: %s\n", p1.Swap())

	// Pair of string and float64
	p2 := Pair[string, float64]{First: "pi", Second: 3.14159}
	fmt.Printf("Original: %s\n", p2)
	fmt.Printf("Swapped: %s\n", p2.Swap())
}

// ============================================================================
// METHOD 7: Generic with Multiple Type Parameters and Constraints
// ============================================================================

// JsonMap converts a slice of T to a map of U (based on a key function)
func SliceToMap[T any, K comparable, V any](
	slice []T,
	keyFunc func(T) K,
	valueFunc func(T) V,
) map[K]V {
	result := make(map[K]V)
	for _, item := range slice {
		result[keyFunc(item)] = valueFunc(item)
	}
	return result
}

func genericMultiParameterExample() {
	fmt.Println("\n--- Method 7: Multiple Type Parameters and Constraints ---")

	type Person struct {
		ID   int
		Name string
		Age  int
	}

	people := []Person{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	// Create a map: ID -> Name
	idToName := SliceToMap(
		people,
		func(p Person) int { return p.ID },
		func(p Person) string { return p.Name },
	)
	fmt.Printf("ID to Name: %v\n", idToName)

	// Create a map: Name -> Age
	nameToAge := SliceToMap(
		people,
		func(p Person) string { return p.Name },
		func(p Person) int { return p.Age },
	)
	fmt.Printf("Name to Age: %v\n", nameToAge)
}

// ============================================================================
// METHOD 8: Generic Interface with Constraints
// ============================================================================

// Sortable constraint: types that support sorting
type Sortable interface {
	int | int32 | int64 | float32 | float64 | string
}

// BubbleSort implements bubble sort for any Sortable type
func BubbleSort[T Sortable](slice []T) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func genericInterfaceExample() {
	fmt.Println("\n--- Method 8: Generic with Interface Constraints ---")

	// Sort integers
	numbers := []int{5, 2, 8, 1, 9, 3}
	BubbleSort(numbers)
	fmt.Printf("Sorted numbers: %v\n", numbers)

	// Sort strings
	strings := []string{"zebra", "apple", "mango", "banana"}
	BubbleSort(strings)
	fmt.Printf("Sorted strings: %v\n", strings)

	// Sort floats
	floats := []float64{3.14, 2.71, 1.41, 1.73}
	BubbleSort(floats)
	fmt.Printf("Sorted floats: %v\n", floats)
}

// ============================================================================
// METHOD 9: Nested Generics
// ============================================================================

// Tree node with generic value
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func (n *TreeNode[T]) Insert(value T, compare func(T, T) int) {
	if compare(value, n.Value) < 0 {
		if n.Left == nil {
			n.Left = &TreeNode[T]{Value: value}
		} else {
			n.Left.Insert(value, compare)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode[T]{Value: value}
		} else {
			n.Right.Insert(value, compare)
		}
	}
}

func (n *TreeNode[T]) InOrder(visit func(T)) {
	if n == nil {
		return
	}
	n.Left.InOrder(visit)
	visit(n.Value)
	n.Right.InOrder(visit)
}

func nestedGenericsExample() {
	fmt.Println("\n--- Method 9: Nested Generics (Binary Search Tree) ---")

	// Tree of integers
	tree := &TreeNode[int]{Value: 5}
	values := []int{3, 7, 2, 4, 6, 8}
	for _, v := range values {
		tree.Insert(v, func(a, b int) int { return a - b })
	}

	fmt.Print("In-order traversal: ")
	tree.InOrder(func(v int) {
		fmt.Printf("%d ", v)
	})
	fmt.Println()
}

// ============================================================================
// METHOD 10: Generic Function with Comparable Constraint
// ============================================================================

// Equal checks if two values are equal (for comparable types)
func Equal[T comparable](a, b T) bool {
	return a == b
}

// RemoveDuplicates removes duplicate elements from a slice
func RemoveDuplicates[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

func comparableConstraintExample() {
	fmt.Println("\n--- Method 10: Comparable Constraint ---")

	fmt.Printf("Equal(42, 42): %v\n", Equal(42, 42))
	fmt.Printf("Equal(42, 43): %v\n", Equal(42, 43))
	fmt.Printf("Equal(\"hello\", \"hello\"): %v\n", Equal("hello", "hello"))

	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	unique := RemoveDuplicates(numbers)
	fmt.Printf("Remove duplicates: %v -> %v\n", numbers, unique)

	words := []string{"go", "rust", "go", "python", "go", "rust"}
	uniqueWords := RemoveDuplicates(words)
	fmt.Printf("Remove duplicates: %v -> %v\n", words, uniqueWords)
}

// ============================================================================
// BONUS: Generic Constraints Cheat Sheet
// ============================================================================

// any - accepts any type
func AcceptAny[T any](value T) T {
	return value
}

// comparable - accepts types that support == and !=
func CheckEquality[T comparable](a, b T) bool {
	return a == b
}

// Numeric-like constraint
type Numeric interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64
}

// Comparable numbers constraint
type ComparableNumeric interface {
	int | int32 | int64 | float32 | float64 | uint | uint32 | uint64 | string
}

// Main demo function
func demoGenerics() {
	fmt.Println("========================================")
	fmt.Println("    GENERICS - All Patterns Demo")
	fmt.Println("========================================")

	basicGenericsExample()
	constrainedGenericsExample()
	genericSliceOperationsExample()
	genericStackExample()
	genericQueueExample()
	genericPairExample()
	genericMultiParameterExample()
	genericInterfaceExample()
	nestedGenericsExample()
	comparableConstraintExample()

	fmt.Println("\n========================================")
	fmt.Println("    GENERICS CONSTRAINTS REFERENCE")
	fmt.Println("========================================")
	fmt.Println("1. 'any'                - Accepts any type")
	fmt.Println("2. 'comparable'         - Supports == and !=")
	fmt.Println("3. 'int|float64|...'    - Union type constraint")
	fmt.Println("4. Custom interfaces    - Define your own constraints")
}

// Uncomment the line below to run this demo
// func main() {
// 	demoGenerics()
// }
