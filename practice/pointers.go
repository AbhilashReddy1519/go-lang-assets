package main

import "fmt"

// pointersDemo shows common pointer patterns and pitfalls in Go.
func pointersDemo() {
    fmt.Println("== basic pointer usage ==")
    x := 10
    p := &x
    fmt.Println("x", x, "*p", *p)
    *p = 20
    fmt.Println("after *p=20 -> x", x)

    fmt.Println("== new vs &var ==")
    n := new(int)
    *n = 5
    fmt.Println("new int ->", *n)
    y := 7
    py := &y
    fmt.Println("&y ->", *py)

    fmt.Println("== nil pointer ==")+    
    var np *int
    fmt.Println("np is nil?", np == nil)

    fmt.Println("== pointer to struct and field mutation ==")
    type Point struct{ X, Y int }
    pt := &Point{X: 1, Y: 2}
    move(pt)
    fmt.Println("moved pt:", pt)

    fmt.Println("== pointer receiver vs value receiver ==")
    c := Counter{v: 1}
    c.IncValue()        // value receiver, does not mutate original
    fmt.Println("after IncValue ->", c.v)
    c.IncPointer()      // pointer receiver, mutates
    fmt.Println("after IncPointer ->", c.v)

    fmt.Println("== pointer to pointer ==")
    a := 3
    pa := &a
    ppa := &pa
    fmt.Println("a, *pa, **ppa ->", a, *pa, **ppa)

    fmt.Println("== return pointer to local (safe) ==")
    pLocal := makeInt()
    fmt.Println("makeInt returned ->", *pLocal)

    fmt.Println("== arrays and pointers ==")
    arr := [3]int{1, 2, 3}
    parr := &arr
    parr[1] = 99
    fmt.Println("arr after parr[1]=99 ->", arr)

    fmt.Println("== nil-interface vs interface-holding-nil pointer ==")
    var i interface{} = (*int)(nil)
    fmt.Println("interface holding nil pointer is nil?", i == nil)
    var j interface{} = nil
    fmt.Println("plain nil interface is nil?", j == nil)

    fmt.Println("== note: no pointer arithmetic in Go ==")
}

func move(p *struct{ X, Y int }) {
    p.X += 10
    p.Y += 10
}

type Counter struct{ v int }

func (c Counter) IncValue() { c.v++ }
func (c *Counter) IncPointer() { c.v++ }

func makeInt() *int {
    v := 42
    return &v
}
