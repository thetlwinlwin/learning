package main

import (
	"fmt"

	"com.example/calculator"
	"rsc.io/quote"
)

func main() {
	result := calculator.Sum(1, 3)
	fmt.Println(result)
	fmt.Println(quote.Hello())
}
