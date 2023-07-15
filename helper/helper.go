package helper

import (
	"strings"
)


func ValidateUserInput(firstName string, lastName string, email string, userTickets int, availableTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 5 && len(lastName) >= 5
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= int(availableTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}