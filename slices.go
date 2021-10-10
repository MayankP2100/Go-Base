package main

import "fmt"

func slicing() {
	x := [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Println(s)
	fmt.Println(s[:cap(s)]) //cap = capacity
	fmt.Println(s[1:])      //Sliced the slice

	fmt.Printf("%T", s) //Type of the array "s"
}
