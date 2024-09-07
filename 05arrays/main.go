package main

import "fmt"

func main() {
	// Declare an array of 5 integers that is initialized with some values
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	// Iterate over the array using the range keyword
	for i, number := range numbers {
		fmt.Println(i, number)
	}

	fmt.Println("Length of array:", len(numbers))
}
