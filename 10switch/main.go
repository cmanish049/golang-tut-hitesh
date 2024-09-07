package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This is a comment
	fmt.Println("Switch and case in Golang")

	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 5
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
	case 3:
		// fallthrough to execute the next case as well
		fmt.Println("Dice number is 3")
		fallthrough
	case 4:
		fmt.Println("Dice number is 4")
	case 5:
		fmt.Println("Dice number is 5")
	default:
		fmt.Println("Dice number is 6")
	}
}
