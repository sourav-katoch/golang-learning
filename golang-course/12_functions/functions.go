package main

// // syntax to declare a function. return data type must be mentioned. function to add two numbers
// // if we have same type of parameter we can specify it only once like add(a, b int) int
// func add(a int, b int) int {
// 	return a + b
// }

// if we have multiple return value we have to specify the data type in the same order as they are returned.
// func getLanguages() (string, string, int) {
// 	return "golang", "javascript", 3
// }

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int{
	return func(a int)int{
		return 1
	}
}

func main() {
	// result := add(3, 5)
	// fmt.Println(result)
	// fmt.Println(getLanguages())

	//we can also get the return value stored in a variable like below
	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)

	//functions in go are first class citizen functions that means we can assign the func to a variable as well
	// we can also pass a func as an argument to another func
	// fn := func(a int) int {
	// 	return 2

	fn := processIt()
	fn(6)
	}
	processIt(fn)

