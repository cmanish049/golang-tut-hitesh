package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// Iterate over the slice using range
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// Iterate over the slice using range
	for i, day := range days {
		fmt.Printf("Day %d of the week is %s\n", i, day)
	}

	// Iterate over the slice using range
	for _, day := range days {
		fmt.Printf("Day of the week is %s\n", day)
	}

	// Iterate over the slice using range
	for i := range days {
		fmt.Printf("Day %d of the week is %s\n", i, days[i])
	}
}
