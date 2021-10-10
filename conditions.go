package main

import (
	"fmt"
)

func cond() {
	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scanln(&age)
	fmt.Println("")

	if age < 18 {
		fmt.Print("Sorry, You are not allowed!")
	} else if age == 18 {
		fmt.Print("Damn, Welcome! You are just on the line")
	} else if age > 18 {
		fmt.Print("Welcome, sir!")
	} else {
		fmt.Printf("%d is not a valid age, please use a number", age)
	}
}
