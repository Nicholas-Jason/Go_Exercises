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

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to our TodoList App!")
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	const greeting = "Hello user. Welcome to our Todolist App!"
	fmt.Fprintf(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Enter Task you wish to ask")

	//var task string
	//fmt.Fscanln(reader, &task)
	//taskItems = append(taskItems, task)
}
func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/add-tasks", addTask)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)

}
