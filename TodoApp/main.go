package main

import (
	"fmt"
	"net/http"
)

var taskOne = "Practice Go"
var taskTwo = "Workout"
var taskThree = "Referesh my skills with Node JS"
var taskFour = "Take a Break"
var taskFive = "Go job hunting"

var taskItems = []string{taskOne, taskTwo, taskThree, taskFour, taskFive}

//add more functions

func showTask(taskItems []string) {
	fmt.Println("List of my Todos")

	for index, task := range taskItems {
		fmt.Println(index+1, ":", task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTasks = append(taskItems, newTask)
	return updatedTasks
}
func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintf(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, taskItems)
}
func main() {
	fmt.Println("###### Welcome to our TodoList App! ######")

	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Do you wish to update TodoList (Y/N)")
	var answer string
	fmt.Scan(&answer)
	if answer == "Y" {
		var newTask string
		fmt.Println("Add task:")
		fmt.Scan(&newTask)
		taskItems = addTask(taskItems, newTask)
		showTask(taskItems)
	} else {
		showTask(taskItems)
	}

}
