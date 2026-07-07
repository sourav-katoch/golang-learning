package main

import "fmt"

//array is a numbered sequence of specific length

func main() {
	// declaration syntax
	var nums [4]int

	nums[0] = 1

	//function to get the length of the array
	fmt.Println(len(nums))

	// function to get a specific value at an index in an array
	fmt.Println(nums[0])

	fmt.Println(nums)

	//arrays have a 0 or false or blANK value by default called as zeroed values. int is zero string is empty string, boolean is empty.
	var vals [4]bool
	fmt.Println(vals)

	var names [3]string
	names[0] = "Sourav"
	fmt.Println(names)

	//another/shorter way to declare elements in an array
	numss := [3]int{1, 2, 3}
	fmt.Println(numss)

	//2d array
	//Generally, slices are used not arrays. use arrays when data has a fixed size, predictable, memory optimization.
	twodee := [2][2]int{{2, 4}, {5, 6}}
	fmt.Println(twodee)

}
