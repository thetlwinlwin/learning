package slicing

import "fmt"

// size is dynamic

// three components
// pointer to 1st reachable ele (not necessarily array's 1st element)
// length of slice
// capacity of slice

func Slicing() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
}

// [January February March] 3 12  	# total of 12 elem start from 0 in original
// [April May June] 3 9				# total of 9 elem left start from 3 in original
// [July August September] 3 6		# total of 6 elem left start from 6 in original
// [October November December] 3 3	# total of 3 elem left start from 9 in original

// slice operator
// s[i:p]

// s represents array
// i is pointer to 1st ele (not necessarily s[0])
// p is last ele for slice

// capacity of slice tell you how much you can extend the slice

func ExtendSlicing() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	quarter2 := months[3:6]
	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
}

// [April May June] 3 9
// [April May June July] 4 9 => add one more to the last position

// Append
// when slice does not have enough capacity
// IT DOUBLES the capacity

func AppendEles() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

// Remove

// 1st approach
// use slice operator

func RemoveEles1() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2
	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove =>", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("After", letters)
	}
}

// 2nd approach
// copy
// when you change an ele from a slice, you're changing the underlying array too.
// Any other slices that refer to the same underlying array will be affected

func EffectOnUnderlyingArr() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
}

func RemoveEles2() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])
	slice1[1] = "Z"

	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)

}
