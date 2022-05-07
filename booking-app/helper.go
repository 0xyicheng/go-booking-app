package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicket bool = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicket
}
