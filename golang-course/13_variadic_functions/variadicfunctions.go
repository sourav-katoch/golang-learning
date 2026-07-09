package main

import "fmt"

// variadic functions accept any number of arguments. Syntax is below
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	// can also pass slice in variadic function
	nums := []int{3, 4, 5, 6}
	result := sum(nums...)
	fmt.Println(result)
}
