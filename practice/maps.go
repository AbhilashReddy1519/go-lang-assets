package main

import (
	"fmt"
	"sync"
)

// mapsDemo demonstrates common map usage patterns in Go.
func mapsDemo() {
	fmt.Println("== map literals and make ==")
	// map literal
	scores := map[string]int{"alice": 90, "bob": 85}
	fmt.Println("scores:", scores)

	// make with initial capacity
	cache := make(map[string]string, 10)
	cache["k1"] = "v1"
	fmt.Println("cache k1:", cache["k1"]) // access

	fmt.Println("== ok idiom (existence check) ==")
	v, ok := scores["carol"]
	fmt.Println("carol value:", v, "exists?", ok)

	fmt.Println("== delete from map ==")
	delete(scores, "bob")
	fmt.Println("scores after delete:", scores)

	fmt.Println("== iterate and map length ==")
	for name, sc := range scores {
		fmt.Printf("%s => %d\n", name, sc)
	}
	fmt.Println("len(scores)", len(scores))

	fmt.Println("== maps as sets ==")
	set := map[string]struct{}{}
	set["read"] = struct{}{}
	set["write"] = struct{}{}
	if _, ok := set["read"]; ok {
		fmt.Println("read in set")
	}

	fmt.Println("== nested maps and maps of slices ==")
	nested := map[string]map[string]int{}
	nested["group1"] = map[string]int{"a": 1}
	nested["group1"]["b"] = 2
	fmt.Println("nested:", nested)

	mpOfSlices := map[string][]int{}
	mpOfSlices["evens"] = []int{2, 4, 6}
	mpOfSlices["evens"] = append(mpOfSlices["evens"], 8)
	fmt.Println("mpOfSlices:", mpOfSlices)

	fmt.Println("== zero value (nil) maps ==")
	var m map[string]int // nil map
	fmt.Println("nil map len", len(m))
	// reading from nil map returns zero value
	fmt.Println("read missing key from nil map =>", m["x"])
	// writing to nil map would panic; create before write
	if m == nil {
		m = make(map[string]int)
	}
	m["ok"] = 1

	fmt.Println("== copying maps (shallow copy) ==")
	src := map[string]string{"a": "1", "b": "2"}
	dst := copyMap(src)
	dst["a"] = "100"
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)

	fmt.Println("== sync.Map for concurrent access ==")
	var sm sync.Map
	sm.Store("k", "v")
	if val, ok := sm.Load("k"); ok {
		fmt.Println("sync.Map k =>", val)
	}

	fmt.Println("== map ordering note ==")
	// map iteration order is randomized; don't rely on iteration order.

	fmt.Println("== passing maps to functions (reference semantics) ==")
	modifyMap(src)
	fmt.Println("after modifyMap src:", src)
}

func copyMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func modifyMap(m map[string]string) {
	m["c"] = "3"
}
