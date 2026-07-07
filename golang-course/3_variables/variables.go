package main

import "fmt"

func main() {
	//declare string
	//has to use the variable, unused variable needs to be deleted
	//syntax var variable_name variable type = "value"
	var name string = "golang"
	fmt.Println(name)
	//does not need to declare type, golang will infer type by itself
	var names = "Go lang will infer type"
	fmt.Println(names)
	//shortand syntax
	shortname := "This is the shortest syntax"
	fmt.Println(shortname)

	//boolean
	var isAdult bool = true
	var is_adult = true
	fmt.Println(isAdult)
	fmt.Println(is_adult)

	//declaring int
	var age int = 18
	fmt.Println(age)
	//shorthand
	shortage := 23
	fmt.Println(shortage)

	//float
	var count float32 = 4.3
	var counts = 5.2
	fmt.Println(count)
	fmt.Println(counts)
	//shorthand
	short_count := 5.4
	fmt.Println(short_count)
}
