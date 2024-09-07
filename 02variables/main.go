package main

import "fmt"

func main() {
	var username string = "admin"
	fmt.Println(username)
	fmt.Printf("Type: %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type: %T\n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Type: %T\n", smallVal)
}
