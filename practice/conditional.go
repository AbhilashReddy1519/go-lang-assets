package main

import "fmt"

func condition() {
	// IF / ELSE
	// Basic conditional execution depending on a boolean expression.
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are not an adult.")
	}

	// IF with short statement
	// You can initialize a variable in the if statement scope.
	if score := 75; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// NOTE: Go does NOT have a ternary operator like C/C++ or Java.
	// Use if/else instead for conditional assignment.
	status := ""
	if age >= 18 {
		status = "adult"
	} else {
		status = "minor"
	}
	fmt.Printf("Status: %s\n", status)

	// SWITCH
	// Switch is often cleaner than long if/else chains for discrete values.
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Unknown day")
	}

	// SWITCH with multiple values in one case.
	// Useful when the same logic applies to several values.
	letter := "e"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant or not a vowel")
	}

	// SWITCH without an expression acts like if/else.
	// This can be useful for range-based conditions.
	number := 42
	switch {
	case number%2 == 0:
		fmt.Println("Even number")
	default:
		fmt.Println("Odd number")
	}

	// FOR LOOPS
	// The only loop keyword in Go is for.
	// Traditional three-part loop.
	for i := 0; i < 3; i++ {
		fmt.Printf("for loop i=%d\n", i)
	}

	// WHILE-LIKE LOOP
	// Omit init and post statements for while-like behavior.
	count := 0
	for count < 3 {
		fmt.Printf("while-like count=%d\n", count)
		count++
	}

	// INFINITE LOOP
	// This loops forever until break is used.
	// Use with caution; break stops the loop.
	for {
		fmt.Println("infinite loop iteration")
		break
	}

	// RANGE LOOP
	// Iterate over arrays, slices, maps, strings, and channels.
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		fmt.Printf("name[%d] = %s\n", index, name)
	}

	// Use underscore to ignore an unused value from range.
	for _, name := range names {
		fmt.Println("Hello", name)
	}

	// BREAK and CONTINUE
	for i := 1; i <= 5; i++ {
		if i == 2 {
			fmt.Println("skip 2")
			continue // skip the rest of this iteration
		}
		if i == 4 {
			fmt.Println("stop at 4")
			break // exit the loop early
		}
		fmt.Println("loop value", i)
	}

	// FALLTHROUGH in switch
	// Use to continue execution to the next case.
	// This is not automatic in Go, unlike C/C++.
	switch 2 {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
	default:
		fmt.Println("default case")
	}

	// SELECT for channels is another conditional form.
	// It waits on multiple channel operations.
	// Here we demonstrate a simple non-blocking select.
	ch := make(chan string, 1)
	ch <- "message"

	select {
	case msg := <-ch:
		fmt.Println("Received from channel:", msg)
	default:
		fmt.Println("No message ready")
	}
}
