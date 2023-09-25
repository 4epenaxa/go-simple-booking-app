package main

import (
	"booking-app/helper"
	"fmt"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// v1 empty list of maps of strings
// var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// v2 empty list of struct
var bookings = make([]UserData, 0)

func main() {
	greetUser()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidEmail, isValidName, isValidTicketsNumber := helper.ValidateUserInput(email, firstName, lastName, userTickets, remainingTickets)
		if isValidEmail && isValidName && isValidTicketsNumber {
			bookTickets(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of all bookings are:\n %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conf is booked out. See you in next year.")
				break // end program
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name must be at least two letters long.")
			}
			if !isValidEmail {
				fmt.Println("Your email must consist of @.")
			}
			if !isValidTicketsNumber {
				if userTickets <= 0 {
					fmt.Printf("You try to buy wrong number of tickets.\n")
				} else {
					fmt.Printf("You try to buy %v and there only %v tickets.\n", userTickets, remainingTickets)
				}
			}
		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking app.\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here\n")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//v1
		//firstNames = append(firstNames, booking["firstNames"])
		//v2
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Print("Enter you first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter you last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter you email: ")
	fmt.Scan(&email)
	fmt.Print("Enter you number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	//v1
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//v2
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of booking:\n %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}
