package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	//fmt.Printf("conferenceTickets is %T,RemainingTickets is %T,ConferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)
			firstNames := getsFirstNames(bookings)

			fmt.Printf("These are our all bookings: %v\n", firstNames)
			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out")

			}
		} else {
			if !isValidName {
				fmt.Println("Your Firstname and Lastname should be more than 2 letters ")
			}
			if !isValidEmail {
				fmt.Println("your Email should contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

		}

		//fmt.Printf("The whole Slice: %v\n", bookings)
		//fmt.Printf("The first value: %v\n", bookings[0])
		//fmt.Printf("The Slice type: %T\n", bookings)
		//fmt.Printf("The lenght of Slice: %v\n", len(bookings))

	}

}
func greetUsers(confName string, conferenceTickets uint, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking app", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ur tickets here")
}
func getsFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidTicketNumber
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	var email string

	//ask user for their name
	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your Email: ")
	fmt.Scan(&email)
	fmt.Println("Enter your tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining from %v", remainingTickets, conferenceName)
}
