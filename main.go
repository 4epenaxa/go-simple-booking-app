package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking app.\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here\n")

	for {
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

		if remainingTickets >= userTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation mail at %v\n", bookings[0], userTickets, email)
			fmt.Printf("%v tickets remaining.\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("This is our all bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conf is booked out.")
				break
			}
		} else {
			fmt.Printf("We have only %v tickets. Try once again please.\n", remainingTickets)
		}
	}
}
