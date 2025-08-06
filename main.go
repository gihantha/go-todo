package main

import "fmt"




func main() {
	
	
	fmt.Println("Welcome to our Todolist App!")

	var taskOne = "Buy groceries"
	var taskTwo = "Clean the house"
	var reward = "Reward yourself with a treat after completing tasks!"

	var taskItems = []string{taskOne, taskTwo, reward}

	// fmt.Println()
	// fmt.Println("This is a simple app to manage your tasks.")
	// fmt.Println("You can add, remove, and view tasks.")
	// fmt.Println("Let's get started!")

	// fmt.Println()
	// fmt.Println("List of tasks:")
	
	// fmt.Println()
	// fmt.Println("Shopping tasks:")
	// fmt.Println("1. " + taskOne)
	// fmt.Println("2. " + reward)

	// fmt.Println()
	// fmt.Println("Household tasks:")
	// fmt.Println("1. " + taskTwo)
	// fmt.Println("2. " + reward)

	// fmt.Println()
	// fmt.Println("Thank you for using our Todolist App!")

	// fmt.Println()


	// for _, task := range taskItems {
	// 	fmt.Println(task)
	// }

	// for index, task := range taskItems {
	// 	// fmt.Println(index+1, "." , task)
	// 	fmt.Printf("Task %d: %s\n", index+1, task)
	// }
	// fmt.Println(taskItems)

	printTasks(taskItems)
	  fmt.Println()
	taskItems = addTask(taskItems, "Walk the dog")
	fmt.Println()


	printTasks(taskItems)

	
}

func printTasks(taskItems []string) {
	fmt.Println("Thank you for using our Todolist App!")

	fmt.Println()

	for index, task := range taskItems {
		fmt.Println("hb")
		fmt.Printf("Task %d: %s\n", index+1, task)
	} 
}

func addTask(taskItems []string, newTask string) []string{
	var updatedTaskItems  =  append(taskItems, newTask)
	return updatedTaskItems
	
}
