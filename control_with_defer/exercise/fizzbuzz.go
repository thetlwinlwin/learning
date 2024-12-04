package exercise

import (
	"fmt"
	"strconv"
)

func fizzBuzz(number int) string {
	switch {
	case number%15 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fizz"
	case number%5 == 0:
		return "Buzz"
	}
	return strconv.Itoa(number)
}

func RunFizzBuzz() {
	for num := 1; num <= 100; num++ {
		fmt.Println(fizzBuzz(num))
	}
}
