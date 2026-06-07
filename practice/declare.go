package main

import "fmt"

func declare() {
	fmt.Println("=== VARIABLE DECLARATION TYPES ===")

	// 1. VAR KEYWORD - used for function-scoped and package-scoped variables
	// - Can be declared without initialization (gets zero value)
	// - Can declare multiple variables at once
	// - Most flexible, works everywhere
	var smsLimit int                           // zero value: 0
	var costPerSms float64                     // zero value: 0.0
	var hasPermission bool                     // zero value: false
	var username string                        // zero value: ""
	var name, age, isActive = "John", 30, true // Multiple declarations with values

	fmt.Println("1. VAR Declaration:")
	fmt.Printf("smsLimit: %v, costPerSms: %v, hasPermission: %v, username: %q\n", smsLimit, costPerSms, hasPermission, username)
	fmt.Printf("name: %q, age: %d, isActive: %v\n\n", name, age, isActive)

	// 2. SHORT ASSIGNMENT (:=) - used inside functions only
	// - Cannot be used at package level
	// - Type is inferred from the right-hand side
	// - Cleaner and more concise syntax
	// - Can reassign only if at least one new variable is introduced
	message := "Hello Go"
	count := 42
	score := 95.5
	isValid := true

	fmt.Println("2. SHORT ASSIGNMENT (:=):")
	fmt.Printf("message: %q, count: %d, score: %f, isValid: %v\n\n", message, count, score, isValid)

	// 3. CONST - used for compile-time constants that cannot change
	// - Must be initialized at declaration
	// - Type is inferred or explicitly specified
	// - Immutable throughout the program
	// - More performant than variables (optimized at compile time)
	const pi = 3.14159
	const maxRetries int = 5
	const appName = "GoApp"

	fmt.Println("3. CONST Declaration:")
	fmt.Printf("pi: %v, maxRetries: %d, appName: %q\n\n", pi, maxRetries, appName)

	// 4. BLANK IDENTIFIER (_) - used when you need to ignore a return value
	// - Useful in loops where index is not needed: for _, val := range slice {}
	// - In function calls where you don't need all return values
	// - Discards the value so Go doesn't complain about unused variables
	_, unused := "Keep this", "Discard this"
	fmt.Println("4. BLANK IDENTIFIER - only keeping:", unused)
	fmt.Println()

	// 5. VAR BLOCK - used to group multiple declarations for readability
	// - Common at package level for related configuration variables
	// - Makes code cleaner and more organized
	var (
		serverPort = 8080
		timeout    = 30
		maxWorkers = 10
		debugMode  = false
	)

	fmt.Println("5. VAR BLOCK:")
	fmt.Printf("serverPort: %d, timeout: %d, maxWorkers: %d, debugMode: %v\n\n", serverPort, timeout, maxWorkers, debugMode)

	// 6. CONST BLOCK - groups multiple constants together
	// - Used for related constants like configuration values
	// - Type is inferred for each const
	const (
		StatusPending   = "pending"
		StatusCompleted = "completed"
		StatusFailed    = "failed"
	)

	fmt.Println("6. CONST BLOCK:")
	fmt.Printf("StatusPending: %q, StatusCompleted: %q, StatusFailed: %q\n\n", StatusPending, StatusCompleted, StatusFailed)

	// 7. TYPE INFERENCE - Go automatically determines the type
	// - Reduces boilerplate code
	// - Type is determined at compile time
	// - Static typing still enforced, just implicit declaration
	implicitInt := 100       // type is int
	implicitFloat := 3.14    // type is float64
	implicitString := "test" // type is string
	implicitBool := true     // type is bool

	fmt.Println("7. TYPE INFERENCE:")
	fmt.Printf("implicitInt: %v (type: int), implicitFloat: %v (type: float64)\n", implicitInt, implicitFloat)
	fmt.Printf("implicitString: %q (type: string), implicitBool: %v (type: bool)\n\n", implicitString, implicitBool)

	fmt.Println("=== FORMAT NOTATIONS (Printf, Sprintf, etc.) ===")
	fmt.Printf("In C++: %%d for int, %%f for float, %%s for string, %%x for hex\n")
	fmt.Println("In Go:  Much more comprehensive and flexible notation system:")

	// General format verbs
	fmt.Println("GENERAL VERBS:")
	fmt.Printf("%%v (any value):              %v\n", 42)
	fmt.Printf("%%T (type of value):          %T\n", 42)

	// Boolean
	fmt.Println("\nBOOLEAN:")
	fmt.Printf("%%v for bool:                 %v\n", true)
	fmt.Printf("%%t for bool:                 %t\n\n", true)

	// Integer
	fmt.Println("INTEGER:")
	fmt.Printf("%%d (decimal):                %d\n", 255)
	fmt.Printf("%%o (octal):                  %o\n", 255)
	fmt.Printf("%%x (hex lowercase):          %x\n", 255)
	fmt.Printf("%%X (hex uppercase):          %X\n", 255)
	fmt.Printf("%%b (binary):                 %b\n", 255)
	fmt.Printf("%%c (rune/character):         %c\n\n", 65)

	// Float
	fmt.Println("FLOAT:")
	fmt.Printf("%%f (decimal, 6 decimals):   %f\n", 3.14159)
	fmt.Printf("%%e (scientific notation):    %e\n", 3.14159)
	fmt.Printf("%%E (scientific uppercase):   %E\n", 3.14159)
	fmt.Printf("%%g (compact notation):       %g\n\n", 3.14159)

	// String and byte
	fmt.Println("STRING & BYTE:")
	fmt.Printf("%%s (string):                 %s\n", "hello")
	fmt.Printf("%%q (quoted string):          %q\n", "hello")
	fmt.Printf("%%x (bytes hex):              %x\n\n", "hello")

	// Pointer
	fmt.Println("POINTER:")
	val := 42
	fmt.Printf("%%p (pointer address):        %p\n\n", &val)

	// Width and precision
	fmt.Println("WIDTH & PRECISION:")
	fmt.Printf("%%5d (width 5):               %5d\n", 42)
	fmt.Printf("%%-5d (left-align width 5):   %-5d|\n", 42)
	fmt.Printf("%%05d (zero-padded width 5):  %05d\n", 42)
	fmt.Printf("%%.2f (2 decimal places):     %.2f\n", 3.14159)
	fmt.Printf("%%8.2f (width 8, 2 decimals): %8.2f\n\n", 3.14159)

	// Complex numbers
	fmt.Println("COMPLEX NUMBERS:")
	c := complex(3, 4)
	fmt.Printf("%%v for complex:              %v\n", c)

}
