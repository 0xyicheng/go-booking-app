package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50 //不变的变量
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers() //Func

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicket := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets) //Func

	if isValidName && isValidEmail && isValidTicket {
		bookTicket(userTickets, firstName, lastName, email) //Func

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames() //Func
		fmt.Printf("The first names of bookings are: %v \n", firstNames)

		var noTickets bool = remainingTickets == 0
		if noTickets {
			fmt.Println("Booked out. Come back next year.  ")
			// break
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
	wg.Wait()
}

var wg = sync.WaitGroup{}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availible\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// creat a map of a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("%v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Print("####################\n")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Print("####################\n")
	wg.Done()
}
