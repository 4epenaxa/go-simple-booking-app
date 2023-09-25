package helper

import (
	"strings"
)

func ValidateUserInput(email string, firstName string, lastName string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidEmail := strings.Contains(email, "@")
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketsNumber := userTickets > 0 && remainingTickets >= userTickets
	return isValidEmail, isValidName, isValidTicketsNumber
}
