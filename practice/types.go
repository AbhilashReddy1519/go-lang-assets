package main

import "fmt"

// Celsius is a custom type based on float64.
// Custom types add domain meaning while still using a built-in type underneath.
type Celsius float64

type Person struct {
	Name string
	Age  int
}

type Greeter func(name string) string

func types() {
	fmt.Println("Hello Abhilash")

	// Boolean type: used for true/false logic and condition checks.
	var active bool = true

	// String type: used for textual data, messages, and identifiers.
	var welcome string = "Go types are powerful"

	// Integer types: used for counters, indexes, sizes, and integer arithmetic.
	var num1 int = 10
	var num2 int8 = -10
	var num3 int16 = 300
	var num4 int32 = 70000
	var num5 int64 = 1234567890

	// Unsigned integer types: used when negative values are invalid.
	var num6 uint = 10
	var num7 uint8 = 255
	var num8 uint16 = 60000
	var num9 uint32 = 4000000000
	var num10 uint64 = 9000000000000000000
	var num11 uintptr = 0x1000

	// byte is an alias for uint8 and is used for raw binary data.
	var b byte = 0x7A

	// rune is an alias for int32 and represents a Unicode code point.
	var r rune = '世'

	// Floating-point types: used for fractional values and scientific calculations.
	var pi float32 = 3.14
	var distance float64 = 12.3456789

	// Complex number types: used in advanced math and signal processing.
	var c1 complex64 = complex(1.2, 3.4)
	var c2 complex128 = complex(5.6, 7.8)

	// Array: fixed-size collection of items. Use when size is known and fixed.
	var arr [3]int = [3]int{1, 2, 3}

	// Slice: dynamic-length collection built on arrays.
	slice := []string{"go", "lang", "types"}

	// Map: key/value lookup container.
	phoneBook := map[string]string{
		"Alice": "+1234567890",
		"Bob":   "+0987654321",
	}

	// Struct: groups related fields into one value.
	person := Person{
		Name: "Abhilash",
		Age:  28,
	}

	// Pointer: holds the memory address of another value.
	var agePtr *int = &person.Age

	// Function type: functions can be passed and stored as values.
	var greeter Greeter = func(name string) string {
		return "Hello, " + name
	}

	// Interface: defines behavior by listing method signatures.
	var printer fmt.Stringer = personStringer{person}

	// Channel: used for safe communication between goroutines.
	messages := make(chan string, 1)
	messages <- "channel message"
	msg := <-messages

	// Custom type based on a built-in type adds semantic meaning.
	var temperature Celsius = 23.5

	fmt.Println(active, welcome)
	fmt.Println(num1, num2, num3, num4, num5)
	fmt.Println(num6, num7, num8, num9, num10, num11)
	fmt.Println(b, r)
	fmt.Println(pi, distance)
	fmt.Println(c1, c2)
	fmt.Println(arr, slice, phoneBook)
	fmt.Println(person, *agePtr)
	fmt.Println(greeter("world"))
	fmt.Println(printer.String())
	fmt.Println(msg)
	fmt.Println(temperature)
}

type personStringer struct {
	person Person
}

func (p personStringer) String() string {
	return fmt.Sprintf("Person{Name:%s, Age:%d}", p.person.Name, p.person.Age)
}
