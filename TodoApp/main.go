package main

import "fmt"

var taskOne = "Practice Go"
var taskTwo = "Workout"
var taskThree = "Referesh my skills with Node JS"
var taskFour = "Take a Break"
var taskFive = "Go job hunting"

var taskItems = []string{taskOne, taskTwo, taskThree, taskFour, taskFive}

func main() {
	fmt.Println("###### Welcome to our TodoList App! ######")

	fmt.Println("List of my Todos")

	for _, task := range taskItems {
		fmt.Println(task)
	}

}
