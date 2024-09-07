package main

import "fmt"

// Create a struct type
type Employee struct {
	firstName, lastName string
	age, salary         int
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User Status:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email:", u.Email)
}

func main() {
	// create an instance of the struct
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// Accessing struct fields using the dot operator
	emp1.age = 26
	emp1.salary = 600
	fmt.Println("Employee 1", emp1)

	// print the struct fields
	fmt.Println("First Name:", emp1.firstName)

	// User struct

	manish := User{"Manish", "manish@gmail.com", true, 25}
	fmt.Println("User:", manish)

	manish.GetStatus()
	manish.NewMail()

	fmt.Println("User email:", manish.Email)
}
