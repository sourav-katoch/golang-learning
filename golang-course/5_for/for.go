package main

import "fmt"

// for -> only construct in go for looping

func main() {
	// ##while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// ##infinite loop
	//for {
	//	println("Hello")
	//}

	//##classic for loop
	for i := 0; i <= 3; i++ {
		if i == 1 {
			//break loop will stop at 1
			//continue will skip the iteration
			continue
		}
		fmt.Println(i)
	}
	// range is used to iterate over a loop for that range.
	for i := range 3 {
		fmt.Println("What", i)
	}
}
