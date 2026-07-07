package main

import (
	"fmt"
	"time"
)

func main() {
	//switch is used in case of complex or longer conditions that is not ideal of if else
	//simple switch

	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("other")
	}

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It'sn a weekend!")
	default:
		fmt.Println("It's a weekend")
	}

	//type switch, declare t only when needed otherwise would work fine.
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer", t)
		case string:
			fmt.Println("String", t)
		case bool:
			fmt.Println("Boolean", t)
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI(55.09)
}
