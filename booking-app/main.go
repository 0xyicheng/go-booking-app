package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

const conferenceTickets = 50 //不变的变量
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers() //Func

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) //Func

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTickets, bookings, firstName, lastName, email) //Func
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := getFirstNames() //Func
			fmt.Printf("The first names of bookings are: %v \n", firstNames)

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
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availible\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking) //strings.Field() 可以根据空格进行分片
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println(("Enter your fist name: "))
	fmt.Scan(&firstName)

	fmt.Println(("Enter your last name: "))
	fmt.Scan(&lastName)

	fmt.Println(("Enter your email: "))
	fmt.Scan(&email)

	fmt.Println(("Enter number of tickets: "))
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, bookings []string, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
