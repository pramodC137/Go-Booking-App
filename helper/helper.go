package helper

import "strings"

func ValidateUserInputs(firstName string, lastName string, email string, userTickets uint, remaniningTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isvalidTicketNumber := userTickets > 0 && userTickets <= remaniningTickets

	return isValidName, isValidEmail, isvalidTicketNumber
}
