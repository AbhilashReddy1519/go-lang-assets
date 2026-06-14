package main

import "fmt"

// Package-level types so methods can be declared at file scope.
type User struct {
	Name  string
	Email string
	Age   int
}

type Address struct {
	Street string
	City   string
	Zip    string
}

type Customer struct {
	User    User
	Address Address
}

type ContactInfo struct {
	Phone string
	City  string
}

type Employee struct {
	User
	ContactInfo // embedded field
	Position    string
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// structsDemo is a single entry point for struct examples.
// Call this from main() to see how structs work in Go.
func structsDemo() {
	fmt.Println("=== BASIC STRUCTS ===")

	// Basic struct instantiation using package-level `User` type.
	// Use this when you want to group related data into one value.
	user := User{Name: "Abhilash", Email: "abhilash@example.com", Age: 28}
	fmt.Printf("user: %+v\n", user)

	// Zero-value struct: fields are set to their default values.
	var emptyUser User
	fmt.Printf("emptyUser: %+v\n", emptyUser)

	// Positional struct literal: only when fields are listed in declaration order.
	user2 := User{"Riya", "riya@example.com", 25}
	fmt.Printf("user2: %+v\n", user2)

	fmt.Println("\n=== NESTED STRUCTS ===")

	// Nested structs let you compose complex types using package-level `Customer` and `Address`.
	customer := Customer{
		User: User{Name: "Nina", Email: "nina@example.com", Age: 30},
		Address: Address{
			Street: "123 Main St",
			City:   "Hyderabad",
			Zip:    "500081",
		},
	}
	fmt.Printf("customer: %+v\n", customer)

	// Access nested fields with dot notation.
	fmt.Println("customer city:", customer.Address.City)

	fmt.Println("\n=== ANONYMOUS STRUCTS ===")

	// Anonymous structs are useful for one-off data shapes or temporary values.
	// They are not declared with a named type.
	anon := struct {
		Title string
		Year  int
	}{
		Title: "Go Examples",
		Year:  2026,
	}
	fmt.Printf("anon: %+v\n", anon)

	fmt.Println("\n=== EMBEDDED STRUCTS ===")

	// Embedding makes the fields and methods of the embedded type available on the outer type.
	// Use the package-level `Employee` and `ContactInfo` types.
	employee := Employee{
		User: User{Name: "Sam", Email: "sam@example.com", Age: 35},
		ContactInfo: ContactInfo{
			Phone: "+91-9876543210",
			City:  "Bangalore",
		},
		Position: "Engineer",
	}

	// Embedded fields are promoted, so we can access them directly.
	fmt.Printf("employee name: %s, city: %s, position: %s\n", employee.Name, employee.City, employee.Position)

	fmt.Println("\n=== METHODS ON STRUCTS ===")

	// Call both value and pointer receiver methods.
	fmt.Println(employee.Greeting())
	employee.AgeUp()
	fmt.Printf("employee age after AgeUp: %d\n", employee.Age)

	fmt.Println("\n=== TAGS AND JSON EXAMPLE ===")

	// Struct tags are metadata used by packages like encoding/json.
	product := Product{ID: 101, Name: "Go Book", Price: 499.99}
	fmt.Printf("product: %+v\n", product)

	fmt.Println("\n=== POINTERS TO STRUCTS ===")

	// Use pointers when you want functions or methods to modify the original struct.
	userPtr := &user
	fmt.Printf("userPtr before: %+v\n", userPtr)
	updateEmail(userPtr, "new_email@example.com")
	fmt.Printf("userPtr after: %+v\n", userPtr)

	fmt.Println("\n=== SLICES OF STRUCTS ===")

	// Slices of structs are useful for collections of records.
	users := []User{
		{Name: "A", Email: "a@example.com", Age: 20},
		{Name: "B", Email: "b@example.com", Age: 22},
	}
	for _, u := range users {
		fmt.Printf("list user: %+v\n", u)
	}

	fmt.Println("\n=== WHEN TO USE EACH STRUCT STYLE ===")
	fmt.Println("- Basic struct: use when grouping related properties into one type.")
	fmt.Println("- Nested struct: use when one object contains another object.")
	fmt.Println("- Anonymous struct: use for temporary values or quick tests.")
	fmt.Println("- Embedded struct: use for reuse and composition; embedded fields are promoted.")
	fmt.Println("- Pointer struct: use when you need to modify the original value or keep copies small.")
}

// Greeting is a value receiver method. Use value receivers when the method doesn't modify the struct.
func (e Employee) Greeting() string {
	return fmt.Sprintf("Hello %s, your position is %s", e.Name, e.Position)
}

// AgeUp is a pointer receiver method. Use pointer receivers when the method modifies the struct.
func (e *Employee) AgeUp() {
	e.Age++
}

// updateEmail changes the Email field of a User.
// Use a pointer parameter to modify the original struct.
func updateEmail(u *User, email string) {
	u.Email = email
}
