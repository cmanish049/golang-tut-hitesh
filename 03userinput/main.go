package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	username, _ := reader.ReadString('\n')
	fmt.Println(username)
	fmt.Printf("Type: %T\n", username)

	fmt.Print("Please rate: ")
	rating, _ := reader.ReadString('\n')
	fmt.Println(rating)
	fmt.Printf("Type: %T\n", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating)
		fmt.Printf("Type: %T\n", numRating)
	}
}
