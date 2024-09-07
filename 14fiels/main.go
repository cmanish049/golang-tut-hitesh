package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "Hello from Go!"

	file, error := os.Create("output.txt")

	if error != nil {
		panic(error)
	}

	length, error := io.WriteString(file, content)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Wrote %d bytes to file\n", length)

	file.Close()

	readFile("output.txt")
}

func readFile(fileName string) {
	file, error := os.Open(fileName)

	if error != nil {
		panic(error)
	}

	// read file
	data := make([]byte, 100)
	count, error := file.Read(data)

	if error != nil {
		panic(error)
	}

	fmt.Printf("Read %d bytes from file\n", count)
	fmt.Println(string(data))
}
