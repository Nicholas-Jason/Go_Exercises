package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var taskOne = "Practice Go"
var taskTwo = "Workout"
var taskThree = "Referesh my skills with Node JS"
var taskFour = "Take a Break"
var taskFive = "Go job hunting"

var taskItems = []string{taskOne, taskTwo, taskThree, taskFour, taskFive}
var user = "Jackson Richards"

//add more functions

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

func addTask(writer http.ResponseWriter, request *http.Request) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(writer, "Enter Task you wish to ask")

	task, _ := reader.ReadString('\n')

	task = strings.TrimSpace(task)
	taskItems = append(taskItems, task)
}
func removeTask(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Enter Task you wish to remove")

	var rtask string
	fmt.Scanln(&rtask)
	var new_taskItems []string
	for _, task := range taskItems {
		if task != rtask {
			new_taskItems = append(new_taskItems, task)
		}
	}
	taskItems = new_taskItems
}
func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/hello-go", helloUser)
	http.HandleFunc("/add-tasks", addTask)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/remove-tasks", removeTask)
	http.ListenAndServe(":8080", nil)

}
