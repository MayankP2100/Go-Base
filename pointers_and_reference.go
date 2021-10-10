package main

import "fmt"

func point_the_refero() {
	x := 1
	y := -56
	fmt.Println(^x, ^y, ^^x) // ^ adds 1 to the value and multiples it with minus and reverse for -ve value
	fmt.Println(&x)          // & shows the reference of x thats the value storage

	z := &x
	fmt.Println(z) //z gets the location of x's storage

	*z = 5 // * is called dereference(pointer) and it gives x the value of z from the location reference
	fmt.Println(x)
}
