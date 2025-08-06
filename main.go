package main

import "fmt"

func main() {
	var taskOne = "Buy groceries"
	var taskTwo = "Clean the house"
	reward := "Reward yourself with a treat after completing tasks!"

	var taskItems = []string{"Watch a movie", "Read a book", "Go for a walk"}
	
	fmt.Println("Welcome to our Todolist App!")

	fmt.Println()
	fmt.Println("This is a simple app to manage your tasks.")
	fmt.Println("You can add, remove, and view tasks.")
	fmt.Println("Let's get started!")

	fmt.Println()
	fmt.Println("List of tasks:")
	
	fmt.Println()
	fmt.Println("Shopping tasks:")
	fmt.Println("1. " + taskOne)
	fmt.Println("2. " + reward)

	fmt.Println()
	fmt.Println("Household tasks:")
	fmt.Println("1. " + taskTwo)
	fmt.Println("2. " + reward)

	fmt.Println()
	fmt.Println("Thank you for using our Todolist App!")

	fmt.Println()

	
}
