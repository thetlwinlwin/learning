package exercise

import "fmt"

func primeFinder(number int) bool {

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	if number != 1 {
		return true
	} else {
		return false
	}
}

func PrimeBetween1and20() {
	fmt.Println("Prime numbers less than 20:")
	for number := 1; number < 20; number++ {
		if primeFinder(number) {
			fmt.Printf("%v ", number)
		}
	}
}
