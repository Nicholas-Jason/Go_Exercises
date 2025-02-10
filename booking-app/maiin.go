package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 30
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have ", conferenceTickets, " tickets and there are ", remainingTickets, " left")
	fmt.Println("Get your tickets here to attend")

}
