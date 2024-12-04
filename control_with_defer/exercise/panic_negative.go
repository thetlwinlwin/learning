package exercise

import "fmt"

func panicNegative() {
	var val int
	for {
		fmt.Print("Enter Number : ")
		fmt.Scanf("%d", &val)

		if val == 0 {
			fmt.Println("0 is neither negative nor positive")
		} else if val < 0 {
			panic("You entered a negative number!")
		} else {
			fmt.Println("You entered:", val)
		}

	}
}

func RunPanicNegative() {
	panicNegative()
}
