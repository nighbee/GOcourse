package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p Person) greeting() {
	fmt.Println("Hello, my name is", p.Name)
}
func main() {
	person := Person{"John", 25}
	person.greeting()
}

/*
How do you define a struct in Go?

    A struct is defined using the type keyword followed by the struct name and the struct
keyword with field definitions inside curly braces.

How do methods differ from regular functions in Go?

    Methods are functions that have a receiver argument, which allows the function
to be associated with a type (e.g., a struct). Methods can access the fields of the struct they are associated with.

Can a method in Go be associated with types other than structs?

    Yes, methods can be associated with any type, including user-defined types, pointers, and even built-in types.
*/
