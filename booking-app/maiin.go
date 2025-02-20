package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 30
var bookings = []string{}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and there are %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstName() {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are all the bookings %v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicket = userTickets > 0 && userTickets < remainingTickets
	return isValidName, isValidEmail, isValidTicket
}
func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}
func bookConf(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you User %v %v for booking %v tickets. We shall send a confirmation email to the address %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v remaining tickets for %v \n", remainingTickets, conferenceName)
}
func main() {

	greetUsers()

	for len(bookings) < 50 && remainingTickets > 0 {
		var firstName, lastName, email, userTickets = getUserInput()

		var isValidName, isValidEmail, isValidTicket = validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicket {
			bookConf(userTickets, firstName, lastName, email)
			//call function print firs names
			printFirstName()
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is completely booked")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Println("Invalid Email. Try again")
			}
			if !isValidTicket {
				fmt.Println("Invalid Ticket Number. Try again")
			}
			if !isValidName {
				fmt.Println("Invalid Name. Try again")
			}
		}

	}

}
