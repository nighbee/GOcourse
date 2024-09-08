package main

func add(a, b int) int {
	return a + b
}
func swap(a, b int) (int, int) {
	return b, a
}
func divide(a, b int) (int, int) {
	return a / b, a % b
}

/*
defo func wth multiple returns:func name(params) (type1, type2) {}
What is the significance of named return values in Go?

    Named return values act as variables that are initialized when the function begins,
which can make the code more readable.

How can you ignore certain return values if you don't need them?

    Use an underscore _ to ignore unwanted return values
*/

func main() {

}
