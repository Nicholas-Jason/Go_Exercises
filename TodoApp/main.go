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
var user = "Jackson Richards"

type Post struct {
	Task string `json: task`
}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome to our TodoList App!")
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello " + user + ". This is your personal Todolist!"
	fmt.Fprintf(writer, "%s", greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for index, task := range taskItems {
		var s string = fmt.Sprint(index)
		fmt.Fprintln(writer, s+" "+task)
	}
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)

}
