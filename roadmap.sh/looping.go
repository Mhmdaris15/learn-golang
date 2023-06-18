package main

import "fmt"

func main() {
	// Print the numbers 1 through 10
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// create variable named x and assign it the value 5 with int8
	var x int8 = 5
	for x >= 0 {
		fmt.Print(x, " ")
		x--
	}

}
