package main

import (
	"fmt"
	"io"
	"strings"
)

// Basic interface example: implicit implementation
// Use the existing `Person` type from `types.go` and `functions.go`.
type Greetable interface{ Greet() string }

func basicInterface() {
	var g Greetable = Person{Name: "Alice", Age: 30}
	fmt.Println(g.Greet())
	sayHello(g)
}

func sayHello(g Greetable) { fmt.Println("sayHello:", g.Greet()) }

// Pointer receiver vs value receiver demonstration and nil interface pitfall
type PointerGreeter struct{ Name string }

func (p *PointerGreeter) Greet() string {
	if p == nil {
		return "<nil>"
	}
	return "Hi, " + p.Name
}

func pointerReceiverDemo() {
	var g Greetable
	pg := &PointerGreeter{Name: "Bob"}
	g = pg
	fmt.Println(g.Greet())

	var nilPg *PointerGreeter = nil
	var g2 Greetable = nilPg
	fmt.Println("interface with nil underlying value equal nil? ->", g2 == nil)
	if g2 == nil {
		fmt.Println("g2 is nil")
	} else {
		fmt.Println("g2 is non-nil (holds a type but nil value)")
	}
}

// Empty interface and type switches/assertions
func emptyInterfaceDemo() {
	describe(42)
	describe("foo")
	describe([]int{1, 2, 3})
}

func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Printf("type %T value %v\n", v, v)
	}
}

// Using interfaces from the standard library: io.Reader
func readerDemo() {
	r := strings.NewReader("hello reader")
	buf := make([]byte, 5)
	n, _ := r.Read(buf)
	fmt.Println(string(buf[:n]))

	printFromReader(strings.NewReader("using io.Reader"))
}

func printFromReader(r io.Reader) {
	b := make([]byte, 1024)
	n, _ := r.Read(b)
	fmt.Println(string(b[:n]))
}

// Interface embedding
type Speaker interface{ Speak() string }
type Mover interface{ Move() string }
type Robot interface {
	Speaker
	Mover
}

type Bot struct{}

func (Bot) Speak() string { return "beep" }
func (Bot) Move() string  { return "rolling" }

func embeddingDemo() {
	var r Robot = Bot{}
	fmt.Println(r.Speak(), "and", r.Move())
}

// Type assertions and switches
func assertionDemo() {
	var i interface{} = "hello"
	s, ok := i.(string)
	fmt.Println(s, ok)

	switch v := i.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	default:
		fmt.Println("unknown", v)
	}
}

// Returning an interface from a constructor
type Counter interface {
	Inc()
	Value() int
}

type counter struct{ v int }

func (c *counter) Inc()       { c.v++ }
func (c *counter) Value() int { return c.v }

func NewCounter() Counter { return &counter{} }

func returnInterfaceDemo() {
	c := NewCounter()
	c.Inc()
	c.Inc()
	fmt.Println("counter:", c.Value())
}

// Public helper to run all interface demos from this package
func interfacesDemo() {
	fmt.Println("== basicInterface ==")
	basicInterface()
	fmt.Println("== pointerReceiverDemo ==")
	pointerReceiverDemo()
	fmt.Println("== emptyInterfaceDemo ==")
	emptyInterfaceDemo()
	fmt.Println("== readerDemo ==")
	readerDemo()
	fmt.Println("== embeddingDemo ==")
	embeddingDemo()
	fmt.Println("== assertionDemo ==")
	assertionDemo()
	fmt.Println("== returnInterfaceDemo ==")
	returnInterfaceDemo()
}
