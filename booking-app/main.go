package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 30
var bookings = make([]map[string]string, 0)

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and there are %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstName() {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	fmt.Printf("These are all the bookings %v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicket = userTickets > 0 && userTickets <= remainingTickets
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
	//create a map for user

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["tickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thank you User %v %v for booking %v tickets. We shall send a confirmation email to the address %v \n", firstName, lastName, userTickets, email)

	fmt.Printf("There are %v remaining tickets for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending ticket:\n %v to email adress %v\n", ticket, email)
	fmt.Println("##############")
}
func main() {

	greetUsers()

	for len(bookings) < 50 && remainingTickets > 0 {
		var firstName, lastName, email, userTickets = getUserInput()

		var isValidName, isValidEmail, isValidTicket = validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicket {
			bookConf(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
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
