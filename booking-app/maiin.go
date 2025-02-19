package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 30

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and there are %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings = []string{}
	for len(bookings) < 50 && remainingTickets > 0 {
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

		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		var isValidTicket = userTickets > 0 && userTickets < remainingTickets

		if isValidTicket && isValidName && isValidEmail {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you User %v %v for booking %v tickets. We shall send a confirmation email to the address %v \n", firstName, lastName, userTickets, email)

			fmt.Printf("There are %v remaining tickets for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("These are all the bookings %v\n", firstNames)
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is completely booked")
				break
			}

		} else {
			fmt.Println("Your input data is invalid. Try again")
		}

	}

}
