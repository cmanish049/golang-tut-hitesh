package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "Javascript"
	languages["rb"] = "Ruby"
	languages["py"] = "Python"
	languages["go"] = "Golang"
	languages["ts"] = "Typescript"

	// Display the map
	fmt.Println("Languages:", languages)

	fmt.Println("Js is", languages["js"])
	fmt.Println("Ruby is", languages["rb"])

	delete(languages, "rb")

	fmt.Println("Languages after deleting Ruby:", languages)

	// loops are interesting
	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}
}
