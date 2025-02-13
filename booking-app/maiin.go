package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 30

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and there are %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	var bookings = []string{}
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Theese are all the bookings %v\n", bookings)

	fmt.Printf("Thank you User %v %v for booking %v tickets. We shall send a confirmation email to the address %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v remaining tickets for %v \n", remainingTickets, conferenceName)
}
