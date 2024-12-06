package arr

import "fmt"

// Array

// must declare dtypes and no of elems
// must assign at once outside of the func
// Go init with default data.
// int => 0
// str => "" (empty str)

/* can't do like this
var a[3] int
a[0] =10 */

func Arr() {
	var a [3]int
	a[1] = 10
	fmt.Println(a[0]) // auto assign zero
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	cities := [5]string{"Yangon", "Mandalay", "Berlin", "Bangkok", "Singapore"}
	fmt.Println("cities", cities)

}

// Ellipsis

// when you don't know how many positions
// use [...]

func Ellipsis() {
	asia := [...]string{"China", "Malaysia", "Thailand"}
	fmt.Println(asia)
}

// can specify the val for last position

func EllipsisLast() {
	numbers := [...]int{99: 123}
	fmt.Println(numbers)
}

// Multidimensional Array

func Multidimensional() {
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
	}
	fmt.Println(twoD)
}
