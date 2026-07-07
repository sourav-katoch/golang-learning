package main

import "fmt"

//maps -> hash, objects, dictionaries

func main() {
	//creating maps
	m := make(map[string]string)

	//setting an element

	m["name"] = "Golang"
	m["area"] = "backend"

	//get an element
	fmt.Println(m["name"])
	fmt.Println(m["area"])
	//IMP : if key value does not exist in map then it returns 0 value (0, empty string or false)
	fmt.Println(m["phone"])

}
