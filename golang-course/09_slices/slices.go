package main

import (
	"fmt"
	"slices"
)

// Slices are basically dynamic arrays as they automatically expand if needed.\
//  We do not need to specify a size.
// most used construct in GO.
//it has useful methods to modify arrays.

func main() {
	//uninitialized slice is nil. If we do not assign values it gets nil by default.
	var nums []int

	//we can also check if it is nil or not as shown below
	fmt.Println(nums == nil)

	//we can check the length of a array/slice as below.
	fmt.Println(len(nums))

	//if we dont want the array to be nil by default we can use the make func as below

	var numb = make([]int, 2, 5)
	fmt.Println(numb)
	fmt.Println(numb == nil)

	//we can check capacity of slice/array which means the maximum number of elements that can be fit.
	//But it gets automatically resized because of it sdynamic nature
	//we can check the capacity using the cap func as below
	//we can add the initial capacity by giving 3rd parameter qwhen declaring a slice.

	fmt.Println(cap(numb))

	//append function to add/append at the end of the array
	numb = append(numb, 1)
	fmt.Println(numb)
	fmt.Println(cap(numb))

	//capacity increases in 2x. if initial capacity is 5 then once it is reached it gets increased to 10.
	//Initially, we keep the lenght and capacity as zero.
	numbs := []int{}
	numbs = append(numbs, 1)
	numbs = append(numbs, 2)
	fmt.Println(numbs)

	var numss = make([]int, 0, 5)
	numss = append(numss, 2)

	var numss2 = make([]int, len(numss))
	fmt.Println(numss, numss2)

	//copy function

	copy(numss2, numss)
	fmt.Println(numss, numss2)

	//slice operator, if we omit the first parameter to slice operator it takes 0 as default.
	// same with the to value in slice operator. It takes the last index as default.

	var count = []int{1, 2, 3}
	fmt.Println(count[0:2])

	//slice package.

	var counts1 = []int{1, 2}
	var counts2 = []int{1, 2}
	//to check if slices are equal.
	fmt.Println(slices.Equal(counts1, counts2))

	//2d slices

	var cnt = [][]int{{1, 2}, {3, 4}}
	fmt.Println(cnt)
}
