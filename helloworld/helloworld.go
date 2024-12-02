package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	firstName, lastName := "Thet", "Lwin"
	age := 28
	fmt.Println(firstName, lastName, age)

	// INT

	// int usually depends on OS bits, has 8,16,32,64
	// uint unsigned number, also have 8,16,32,64
	// signed integer => int is for positive and negative
	// un-signed integer => uint is for positive only

	var integer16 int16 = 8
	var integer8 int8 = 8
	// cannot use math operator on differnet int type
	// fmt.Println(integer16 + integer8)
	fmt.Println(integer16 + int16(integer8))

	// rune is used to represent unicode character
	// have to write with SINGLE QUOTE 'G'
	rune := 'G'
	fmt.Println(rune) // 71

	// overflow error
	// coz it is unsigned
	// var my uint = -1
	// coz it is more than 32 bits
	// var my2 int32 = 234154134583

	// FLOAT

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	const e = 2.71818
	const avogadro = 6.0221419e23
	const planck = 6.626e-34

	// String

	// has to use double quotation marks(")
	// single quotation marks (') are used for single characters (and for runes)
	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)

	// Default

	var defaultInt int         //0
	var defaultFloat32 float32 // 0
	var defaultFloat64 float64 // 0
	var defaultBool bool       // false
	var defaultString string   // empty string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// Conversion

	// Atoi convert string to int
	ii, _ := strconv.Atoi("-42")

	// Itoa convert int to string
	s := strconv.Itoa(-42)
	fmt.Println(ii, s)

}
