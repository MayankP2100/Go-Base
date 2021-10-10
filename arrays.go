package main

import "fmt"

func arraysFunc() {
	var array [5]int
	array[0] = 100
	array[3] = 25
	fmt.Println(array, "\n", array[3])

	arr := [3]int{45, 63, 12}
	fmt.Println(arr)
	fmt.Println(len(arr))
}
