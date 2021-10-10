package main

import "fmt"

func round_n_round() {
	defer fmt.Println("done looping")
	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}
}
