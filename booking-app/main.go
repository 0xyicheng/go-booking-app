package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const confreenceTickets = 50 //不变的变量
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking applicatio.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availible\n", confreenceTickets, remainingTickets)

	for {

		// var bookings []string         //array

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// 输入用户名
		fmt.Println(("Enter your fist name: "))
		fmt.Scan(&firstName)

		fmt.Println(("Enter your last name: "))
		fmt.Scan(&lastName)

		fmt.Println(("Enter your email: "))
		fmt.Scan(&email)

		fmt.Println(("Enter number of tickets: "))
		fmt.Scan(&userTickets)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail bool = strings.Contains(email, "@")
		var isValidTicket bool = userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicket {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking) //strings.Field() 可以根据空格进行分片
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			var noTickets bool = remainingTickets == 0
			if noTickets {
				fmt.Println("Booked out. Come back next year.  ")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("name you entered is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("email you entered is not right\n")
			}
			if !isValidTicket {
				fmt.Printf("number of tickets you entered is invalid\n")
			}
		}
	}

	// Switch statement example
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")
	}

}
