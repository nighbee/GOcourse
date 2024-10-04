package main

import (
	"github.com/nighbee/math1/add"
	"github.com/nighbee/math1/multiply"
	"github.com/nighbee/math1/substract"
)

func main() {
	a, b := 5, 3
	println(add.Add(a, b))
	println(substract.Substract(a, b))
	println(multiply.Multiply(a, b))
}
