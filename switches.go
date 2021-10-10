package main

import "fmt"

func switcheroo() {
	ans := 4
	switch ans {
	case 1:
		fmt.Println("1. one")
		fmt.Println("2, one")
	case 2:
		fmt.Print("two")
	default:
		fmt.Print("not a case")
	}
}
