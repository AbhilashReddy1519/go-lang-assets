# Go Fundamentals - Complete Developer Guide

# Table of Contents

1. Variables
2. Data Types
3. Functions
4. Structs
5. Arrays
6. Slices
7. Maps
8. Pointers
9. Interfaces
10. Error Handling

---

# Variables

Variables store data in memory.

## Declaration

```go
var age int
var name string
```

Multiple declarations:

```go
var (
    age  int
    name string
)
```

---

## Initialization

```go
var age int = 25
```

---

## Type Inference

Go can infer types.

```go
var age = 25
var price = 19.99
var active = true
```

Compiler determines type automatically.

---

## Short Declaration

Only inside functions.

```go
age := 25
name := "John"
```

Cannot be used at package level.

---

## Zero Values

Uninitialized variables receive zero values.

```go
var i int        // 0
var f float64    // 0
var b bool       // false
var s string     // ""
```

---

## Scope

### Package Scope

```go
package main

var appName = "MyApp"
```

Accessible throughout package.

---

### Function Scope

```go
func main() {
    age := 25
}
```

---

### Block Scope

```go
if true {
    x := 10
}
```

x only exists inside block.

---

## Shadowing

Avoid accidental shadowing.

```go
x := 10

if true {
    x := 20
}
```

Different variables.

---

## Constants

Immutable values.

```go
const Pi = 3.14159
```

Multiple constants:

```go
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

---

## Untyped Constants

```go
const x = 100
```

Flexible until assigned.

---

## iota

Auto-incrementing constant generator.

```go
const (
    Sunday = iota
    Monday
    Tuesday
)
```

Output:

```text
0
1
2
```

---

### Bit Masks with iota

```go
const (
    Read = 1 << iota
    Write
    Execute
)
```

Values:

```text
1
2
4
```

---

# Data Types

---

## Integer Types

```go
int
int8
int16
int32
int64

uint
uint8
uint16
uint32
uint64
uintptr
```

Example:

```go
var age int = 25
```

---

## int vs int64

```go
var a int
var b int64
```

Cannot mix directly.

```go
int64(a)
```

---

## Floating Point

```go
float32
float64
```

Example:

```go
price := 99.99
```

---

## Scientific Notation

```go
x := 1.2e6
```

---

## bool

```go
var active bool = true
```

Only:

```go
true
false
```

No truthy/falsy values.

---

## string

Immutable UTF-8 sequence.

```go
name := "Go"
```

---

### Raw Strings

```go
msg := `
Line 1
Line 2
`
```

---

### String Length

```go
len("hello")
```

Returns bytes.

---

## rune

Represents Unicode code point.

Alias for:

```go
int32
```

Example:

```go
var r rune = 'A'
```

---

### Unicode

```go
r := '😊'
fmt.Printf("%c", r)
```

---

## byte

Alias for:

```go
uint8
```

Example:

```go
var b byte = 'A'
```

Useful for binary data.

---

# Functions

---

## Basic Function

```go
func greet() {
    fmt.Println("Hello")
}
```

---

## Parameters

```go
func add(a int, b int) int {
    return a + b
}
```

Short form:

```go
func add(a, b int) int {
    return a + b
}
```

---

## Multiple Returns

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }

    return a / b, nil
}
```

---

## Named Returns

```go
func rectangle(w, h int) (area int) {
    area = w * h
    return
}
```

---

## Variadic Functions

```go
func sum(nums ...int) int {
    total := 0

    for _, n := range nums {
        total += n
    }

    return total
}
```

Usage:

```go
sum(1,2,3,4)
```

---

## Anonymous Functions

```go
func() {
    fmt.Println("Hello")
}()
```

---

## Closures

Closure remembers outer variables.

```go
func counter() func() int {
    count := 0

    return func() int {
        count++
        return count
    }
}
```

---

## First-Class Functions

```go
func add(a,b int) int {
    return a+b
}

var fn func(int,int) int = add
```

---

## Defer

Executes before function returns.

```go
func main() {
    defer fmt.Println("Last")
    fmt.Println("First")
}
```

Output:

```text
First
Last
```

---

### Multiple Defers

LIFO order.

```go
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
```

Output:

```text
3
2
1
```

---

## Panic

Stops normal execution.

```go
panic("something bad happened")
```

---

## Recover

Used inside deferred functions.

```go
func safe() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    panic("boom")
}
```

---

# Structs

Structs group related fields.

---

## Definition

```go
type User struct {
    Name string
    Age  int
}
```

---

## Creation

```go
u := User{
    Name: "John",
    Age:  25,
}
```

---

## Anonymous Struct

```go
person := struct {
    Name string
}{
    Name: "John",
}
```

---

## Methods

```go
func (u User) Greet() {
    fmt.Println("Hello", u.Name)
}
```

---

## Value Receiver

Copies struct.

```go
func (u User) ChangeName() {
    u.Name = "New"
}
```

Original unchanged.

---

## Pointer Receiver

```go
func (u *User) ChangeName() {
    u.Name = "New"
}
```

Modifies original.

---

## Embedding

```go
type Address struct {
    City string
}

type User struct {
    Name string
    Address
}
```

Usage:

```go
u.City
```

---

## Composition

Preferred over inheritance.

```go
type Engine struct {
    Power int
}

type Car struct {
    Engine Engine
}
```

---

## Struct Tags

Used for JSON and reflection.

```go
type User struct {
    Name string `json:"name"`
}
```

---

# Arrays

Fixed-size collections.

---

## Declaration

```go
var nums [5]int
```

---

## Initialization

```go
nums := [5]int{1,2,3,4,5}
```

---

## Compiler Count

```go
nums := [...]int{1,2,3}
```

---

## Memory Layout

Stored contiguously.

```text
[1][2][3][4]
```

Fast index access O(1).

---

## Array Copy

Assignment copies entire array.

```go
a := [3]int{1,2,3}
b := a
```

---

# Slices

Most used collection type.

---

## Creation

```go
nums := []int{1,2,3}
```

---

## make()

```go
nums := make([]int, 5)
```

Length = 5

---

## Length

```go
len(nums)
```

---

## Capacity

```go
cap(nums)
```

---

## Slice From Array

```go
arr := [5]int{1,2,3,4,5}

s := arr[1:4]
```

---

## append

```go
nums = append(nums, 4)
```

Multiple:

```go
nums = append(nums, 4,5,6)
```

---

## Append Another Slice

```go
a := []int{1,2}
b := []int{3,4}

a = append(a, b...)
```

---

## copy

```go
src := []int{1,2,3}

dst := make([]int, len(src))

copy(dst, src)
```

---

## Internal Structure

Slice contains:

```go
type slice struct {
    ptr *Element
    len int
    cap int
}
```

---

## Slice Sharing

```go
a := []int{1,2,3,4}

b := a[:2]

b[0] = 100
```

Both affected.

---

## Hidden Reallocation

When capacity exceeded:

```go
append()
```

New backing array created.

Important for performance.

---

# Maps

Hash table implementation.

---

## Creation

```go
m := make(map[string]int)
```

---

## Literal

```go
m := map[string]int{
    "john": 25,
}
```

---

## Insert

```go
m["alice"] = 30
```

---

## Lookup

```go
age := m["alice"]
```

---

## Existence Check

```go
age, ok := m["alice"]
```

---

## Delete

```go
delete(m, "alice")
```

---

## Iteration

```go
for k, v := range m {
    fmt.Println(k, v)
}
```

Order not guaranteed.

---

## Internal Details

Go maps use:

- Hashing
- Buckets
- Overflow buckets
- Automatic resizing

Average complexity:

```text
Lookup  O(1)
Insert  O(1)
Delete  O(1)
```

---

# Pointers

Store memory addresses.

---

## Address Operator

```go
x := 10

p := &x
```

---

## Dereference

```go
fmt.Println(*p)
```

---

## Modify Through Pointer

```go
*p = 20
```

---

## Function Example

```go
func increment(n *int) {
    *n++
}
```

Usage:

```go
increment(&x)
```

---

## Pointer to Struct

```go
user := &User{
    Name: "John",
}
```

Access:

```go
user.Name
```

Go auto-dereferences.

---

## Heap vs Stack

### Stack

Fast allocation.

```go
func test() {
    x := 10
}
```

---

### Heap

Allocated when variable escapes.

```go
func create() *int {
    x := 10
    return &x
}
```

Compiler moves x to heap.

---

## Escape Analysis

Compiler decides:

```text
Stack
or
Heap
```

Automatically.

---

# Interfaces

Interfaces define behavior.

---

## Definition

```go
type Speaker interface {
    Speak()
}
```

---

## Implementation

No explicit keyword needed.

```go
type Dog struct{}

func (Dog) Speak() {
    fmt.Println("Woof")
}
```

Dog automatically implements interface.

---

## Polymorphism

```go
func makeSound(s Speaker) {
    s.Speak()
}
```

---

## Empty Interface

```go
interface{}
```

Equivalent:

```go
any
```

Can hold any value.

---

## Type Assertions

```go
var x interface{} = "hello"

s := x.(string)
```

Safe version:

```go
s, ok := x.(string)
```

---

## Type Switch

```go
switch v := x.(type) {
case string:
    fmt.Println(v)
case int:
    fmt.Println(v)
}
```

---

## Dependency Injection

```go
type Logger interface {
    Log(string)
}
```

Inject implementation:

```go
type Service struct {
    logger Logger
}
```

Enables testing.

---

## Interface Internals

Contains:

```text
Type Information
+
Data Pointer
```

---

## Nil Interface Trap

```go
var d *Dog = nil

var s Speaker = d

s != nil
```

Common interview question.

---

# Error Handling

Go prefers explicit errors.

---

## error Interface

```go
type error interface {
    Error() string
}
```

---

## Creating Errors

```go
errors.New("something failed")
```

---

## Custom Errors

```go
type ValidationError struct {
    Message string
}

func (e ValidationError) Error() string {
    return e.Message
}
```

---

## Returning Errors

```go
func validate(age int) error {
    if age < 18 {
        return errors.New("underage")
    }

    return nil
}
```

---

## Wrapping Errors

```go
fmt.Errorf("database failed: %w", err)
```

---

## Unwrap

```go
errors.Unwrap(err)
```

---

## errors.Is

Compare wrapped errors.

```go
if errors.Is(err, ErrNotFound) {
}
```

---

## errors.As

Extract specific error type.

```go
var vErr ValidationError

if errors.As(err, &vErr) {
}
```

---

## Sentinel Errors

Predefined reusable errors.

```go
var ErrNotFound = errors.New("not found")
```

Usage:

```go
if err == ErrNotFound {
}
```

Better:

```go
errors.Is(err, ErrNotFound)
```

---

# Important Hidden Go Features

## Blank Identifier

Ignore values.

```go
_, err := os.Open("file.txt")
```

---

## init Function

Runs automatically.

```go
func init() {
    fmt.Println("setup")
}
```

Runs before main.

---

## Multiple init Functions

Allowed.

```go
func init() {}
func init() {}
```

Executed in order.

---

## Range

```go
for i, v := range nums {
}
```

Works on:

- Arrays
- Slices
- Maps
- Strings
- Channels

---

## UTF-8 String Iteration

```go
for _, r := range "Hello 😊" {
    fmt.Printf("%c\n", r)
}
```

Returns runes.

---

## make vs new

### make

For:

```go
slice
map
channel
```

Example:

```go
make([]int, 10)
```

---

### new

Allocates zero value.

```go
p := new(int)
```

Returns:

```go
*int
```

---

## Garbage Collection

Automatic memory cleanup.

No manual:

```text
free()
delete()
```

required.

---

## Concurrency Safety Warning

Maps are NOT safe for concurrent writes.

Need:

```go
sync.Mutex
```

or

```go
sync.Map
```

---

# Common Interview Questions

### Array vs Slice

| Array | Slice |
|---------|---------|
| Fixed Size | Dynamic |
| Value Type | Reference-like |
| Copies Entire Data | Shares Backing Array |

---

### Value Receiver vs Pointer Receiver

| Value | Pointer |
|---------|---------|
| Copy | Original |
| Small Structs | Large Structs |
| Read Only | Modify State |

---

### make vs new

| make | new |
|---------|---------|
| slice/map/channel | Any Type |
| Initialized | Zero Value |
| Returns Value | Returns Pointer |

---

### Panic vs Error

| Error | Panic |
|---------|---------|
| Expected Failure | Unexpected Failure |
| Returned | Crashes Flow |
| Recoverable | Usually Fatal |

---

# Best Practices

1. Prefer slices over arrays.
2. Prefer composition over inheritance.
3. Return errors, don't panic.
4. Keep interfaces small.
5. Use pointer receivers for large structs.
6. Wrap errors using `%w`.
7. Check map lookups with `ok`.
8. Avoid unnecessary heap allocations.
9. Use `go vet` regularly.
10. Write tests for exported functions.

---

# Summary

Master these concepts in order:

1. Variables
2. Types
3. Functions
4. Structs
5. Arrays
6. Slices
7. Maps
8. Pointers
9. Interfaces
10. Error Handling

These form the foundation required before learning:

- Goroutines
- Channels
- Context
- Reflection
- Generics
- Testing
- Concurrency Patterns
- System Design in Go