package main

import "fmt"

// specifying data type is not needed or neccessary, GO will incur
// const value cannot be changed later
// shorthand syntax is not allowed outside function
const name string = "Golang"
const age int = 30

// const can be

const (
	port = 5000
	host = "localhost"
)

func main() {

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(port, host)
}
