package main

import (
	"fmt"
	"net/http"
)

var task1 = "Buy groceries"
var task2 = "Complete homework"
var task3 = "Call mom"

var taskItems = []string{task1, task2, task3}

func main() {

	fmt.Println("Welcome to our Todolist App!")
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)

}

func helloUser(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello, User!"
	fmt.Fprintf(w, greeting)
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintf(w, "%s\n", task)
	}
}

func printTasks(taskItems []string) {
	fmt.Println("Thank you for using our Todolist App!")

	fmt.Println()

	for index, task := range taskItems {
		fmt.Println("hb")
		fmt.Printf("Task %d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	return updatedTaskItems

}
