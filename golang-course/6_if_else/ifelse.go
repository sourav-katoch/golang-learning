package main

import "fmt"

func main() {
	age := 13
	// we can declare a variable inside the if construct
	if age = 20; age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("Is a teenager")
	} else {
		fmt.Println("is a kid.")
	}

	var role = "admin"
	var hasPermissions = false
	// AND &&, OR ||
	if role == "admin" && hasPermissions {
		fmt.Println("Authorized")
	} else {
		fmt.Println("Not Authorized")
	}

	//GO does not have ternary operator

}
