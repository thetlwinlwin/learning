package main

import (
	"fmt"
	"math/rand"
	"time"

	"com.example/calculator"
)

func gimmenumber() int {
	return -1
}

func main() {
	result := calculator.Sum(1, 3)
	if num := result - gimmenumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit left")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// outside of the if scope
	// fmt.Println(num)

	i := rand.Int31n(10)

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	// default is optional
	default:
		fmt.Println("Default")
	}
	fmt.Println("ok")

	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
	invoker()

}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

// switch can also invoke a function
func invoker() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn Go")
	default:
		fmt.Println("Time to rest")
	}
}
