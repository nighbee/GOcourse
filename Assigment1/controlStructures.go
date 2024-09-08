package main

import "fmt"

// Control Structures
// IF ELSE
func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("The number is positive")
	} else if num < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is zero")
	}

	// For
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("The sum of the first 10 numbers is: ", sum)
	// Switch

	var day int
	fmt.Print("Enter a number between 1 and 7: ")
	fmt.Scan(&day)
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
		fmt.Println("Invalid number")
	}
}

// Questions:
/*
	Go requires the condition in the if statement
	but does not require parentheses, and braces {} are mandatory.

	ways of writing for loop:
Traditional: for i := 0; i < 10; i++ {}
As a while: for condition {}
Infinite loop: for {}

in go switch do not require break statement
*/
