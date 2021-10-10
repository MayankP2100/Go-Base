package main

import "fmt"

func data_types() {
	var number = 10          //Golang can guess the variable type on its own and is 99.9% times reliabele
	fmt.Printf("%T", number) //Printf can be used to use format variables, %T shows the type of the var given
	//%t represents boolean values that are true and false

	fmt.Println("") //Just for spacing

	name := "May" //We dont even have to give "var" keyword while using ":="
	fmt.Println(name)

	var age uint8    //It is possible to set variable without giving it an exact value
	fmt.Println(age) //The default value of the variable is always 0

	var standard string
	fmt.Println(standard) //Default value for a string is "", an empty string

	var bl bool
	fmt.Println(bl) //It prints false because the default value of bool is false
	bl = true       //Changing the value to true
	fmt.Println(bl)

	fmt.Printf("1 Number: %b \n", 300)   //%b shows the binary value of the given number
	fmt.Printf("2 Number: %o \n", 300)   //Octal value
	fmt.Printf("3 Number: %d \n", 300)   //Decimal value
	fmt.Printf("4 Number: %x \n\n", 300) //Hexa decimal; %X to get capital letters

	fmt.Printf("5 Number: %e \n", 300.21654123146521)   //Scientific notation
	fmt.Printf("6 Number: %f \n", 300.21654123146521)   //Decimal no eponent; %f and %F both are same
	fmt.Printf("7 Number: %e \n\n", 300.21654123146521) //For large exponents

	fmt.Printf("8 Number: %s \n", "Hello")   //Default for string
	fmt.Printf("9 Number: %q \n\n", "Hello") //Shows string with "double quotes"

	//Padding
	fmt.Printf("10 Number: %05d \n", 56) //Pads digit to x number of characters with 0s
	fmt.Printf("11 Number: %4d \n", 2)   //Pads digit with spaces, do %-X to start padding from left
	// fmt.Printf("11 Number: %9s \n", "Hello") //Can be done with strings too
	// fmt.Printf("11 Number: %08q \n", "Hello")
	fmt.Printf("12 Number: %.3f \n", 300.21654123146521) //Rounds upto given decimal points
}
