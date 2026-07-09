package main

import "fmt"

func main() {

	// num := []int{6, 7, 8}
	// //we can use for loop to iterate over a data structure like below

	// for i := 0; i < len(num); i++ {
	// 	fmt.Println(num[i])
	// }

	// // or we can use Range to iterate. We can create for loop with range function. as below
	// for _, nums := range num {
	// 	fmt.Println(nums)

	// to get sum using range

	// count := []int{5, 7, 1}
	// sum := 0
	// for i, counts := range count {
	// 	sum = sum + counts
	// 	fmt.Println(counts, i)
	// }
	// fmt.Println(sum)

	// we can also iterate over a map

	// m := map[string]string{"fname": "John", "lname": "Doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// // if we only give single value in for then the key will be returned
	// for k1 := range m {
	// // 	fmt.Println(k1)
	// }
	//Range can be used to iterate over a string. It returns starting byte of rune and unicode of the alphabet
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

}
