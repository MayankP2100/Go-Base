package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input_func() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Printf("Your name is %s \n", name)

	var age string
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("and your age is %s", age)

	fmt.Println("")
	ageN, _ := strconv.ParseInt(age, 10, 16)
	fmt.Printf("Your age after 3 years will be %d", ageN+3)

	fmt.Printf("\n\n")

	//CONVERSATION

	fmt.Println("Conversation example")
	var example_of_float_number float32 = 5.91267
	conversing_float_to_int := int(example_of_float_number)
	fmt.Print(conversing_float_to_int)
}
