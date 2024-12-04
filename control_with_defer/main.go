package main

import (
	"fmt"
	"io"
	"os"

	exe "com.example/control_with_defer/exercise"
)

func main() {
	/* 	eg1()
	   	fmt.Println("recovering")
	   	recovering()
	   	fmt.Println("panic")
	   	runHighLow() */
	exe.RunFizzBuzz()
	exe.PrimeBetween1and20()
	exe.RunPanicNegative()
}

// defer

// postpone any function (including params)
// until the function that contains the defer statement finishes.
// last in, first out

// closing files
// running clean up process

func eg1() {
	for i := 0; i <= 4; i++ {
		defer fmt.Println("deferred finish", -i)
		fmt.Println("regular", i)
	}
	fileClosing("lol.txt", "hi Goo")
}

//	regular 1
//	regular 2
//	regular 3
//	regular 4
//	deferred -4
//	deferred -3
//	deferred -2
//	deferred -1

func fileClosing(filename, line string) {
	newfile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error! Could not create file")
		return
	}
	defer newfile.Close()

	if _, error := io.WriteString(newfile, line); error != nil {
		fmt.Println("Error! Could not write data")
		return
	}
}

// panic

// runtime-error makes go program to panic
// this function force the program to panic
// in panic mode, deferred func run smoothly

// panic func, normally send err msg about why u're panicking

func runHighLow() {
	highlow(2, 0)
	// below won't run because program crashed in highlow func call
	fmt.Println("Program finished successfully!")
}

func highlow(high, low int) {
	if high < low {
		fmt.Println("panic")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")
	highlow(high, low+1)
}

// when low > high, the program panics. the control flow is interrupted
// then all the defer function start to print
// from 2 to 0

// recover

// avoid crush and report error internally.
// only call recover in func where you call defer

func recovering() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("recovering (): recover", handler)
		}
	}()
	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}

// there is no stack trace error
// panic recover is like try and catch approach
