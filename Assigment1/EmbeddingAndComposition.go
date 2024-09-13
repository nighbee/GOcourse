package main

import "fmt"

type Employee struct {
	Name string
	Age  int
	ID   int
}
type Manager struct {
	Employee
	TeamSize int
}

func (e Employee) work() {
	fmt.Printf("%s is working\n", e.Name)
}
func main() {
	manager := Manager{Employee{Name: "Alice", Age: 30, ID: 1}, 5}
	manager.work()
}

/*
What is embedding in Go, and how does it relate to composition?

    Embedding is a way to include a struct or other types within another struct, enabling composition. It allows the embedded struct’s fields and methods to be accessed as if they were part of the outer struct.

How does Go handle method calls on embedded types?

    Go allows you to call methods of the embedded type directly from the outer struct, treating them as if they were defined in the outer struct.

Can an embedded type override a method from the outer struct?

    No, but the outer struct can define a method with the same name, which will take precedence when called from the outer struct.
*/
