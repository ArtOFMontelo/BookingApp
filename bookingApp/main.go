package main

import (
	"awesomeProject/helper"
	"fmt"
	"strconv"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	//fmt.Printf("conferenceTickets is %T,RemainingTickets is %T,ConferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			firstNames := getsFirstNames()

			fmt.Printf("These are our all bookings: %v\n", firstNames)
			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out")
				break
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
func greetUsers() {
	fmt.Printf("Welcome to %v booking app", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get ur tickets here")
}
func getsFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings%v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n", firstName, lastName, email)
	fmt.Printf("%v tickets remaining from %v", remainingTickets, conferenceName)
}
