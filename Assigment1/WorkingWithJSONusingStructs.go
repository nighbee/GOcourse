package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func toJSON(p Product) string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}
func fromJSON(s string) Product {
	p := Product{}
	err := json.Unmarshal([]byte(s), &p)
	if err != nil {
		return Product{}
	}
	return p
}
func main() {
	product := Product{Name: "Laptop", Price: 999.99, Quantity: 10}
	jsonStr := toJSON(product)
	fmt.Printf("JSON string: %s\n", jsonStr)
	decodeProduct := fromJSON(jsonStr)
	fmt.Printf("Product: %+v\n", decodeProduct)
}

/*
How do you work with JSON in Go?

    Use the encoding/json package, with json.Marshal to encode and json.Unmarshal to decode JSON data.

What role do struct tags play in JSON encoding/decoding?

    Struct tags define how fields are named when encoded/decoded, helping match JSON keys to struct fields.

How do you handle errors that may occur during JSON encoding/decoding?

    Handle errors by checking the error returned by Marshal or Unmarshal, and respond accordingly, often logging or handling gracefully.
*/
