package main

import "fmt"

func main() {
	// Create a slice of strings
	// with a length of 5 and capacity of 8
	var fruitList = []string{}

	// Display the slice
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	// Display the length of the slice
	fmt.Println("Length of fruitlist is", len(fruitList))

	// Display the capacity of the slice
	fmt.Println("Capacity of fruitlist is", cap(fruitList))

	// append to the slice
	fruitList = append(fruitList, "Apple")
	fruitList = append(fruitList, "Tomato")
	fruitList = append(fruitList, "Banana")
	fruitList = append(fruitList, "Grapes")
	fruitList = append(fruitList, "Cherries")

	// Display the slice
	fmt.Println("Fruit List:", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println("High Scores:", highScores)

	// demo of how to remove an element from a slice

	var courses = []string{"react", "angular", "vue", "svelte", "ember", "nextjs"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("Courses after removing an element:", courses)

}
