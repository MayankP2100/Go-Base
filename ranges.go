package main

import "fmt"

func ranger() {
	a := [4]int{123, 434, 54, 56}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i]) //Printing each element of the array "a"
	}

	fmt.Println("") //Spacing

	for index, value := range a {
		fmt.Printf("%d: %d\n", index, value)
	}
}
