package main

import "fmt"

// Variables and data types
func main() {
	var a int = 10
	var b float64 = 3.14
	var c string = "Hello world!"
	var d bool = true

	x := 10
	y := 3.14
	z := "Hello world!"
	k := true

	// Print the variables

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(k)
}

/*
1) the difference between var and := is that var is used to declare a variable and assign a value to it,
while := is used to declare a variable and assign a value to it without specifying the type of the variable.
2) fmt.Printf("Type: %T\n", a) is used to print the type of the variable a.
3) we cant change the vars type if its already declared.
*/
