package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const confreenceTickets = 50 //不变的变量
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking applicatio.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availible\n", confreenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// 输入用户名
	fmt.Println(("Enter your fist name: "))
	fmt.Scan(&firstName)

	fmt.Println(("Enter your last name: "))
	fmt.Scan(&lastName)

	fmt.Println(("Enter number of tickets: "))
	fmt.Scan(&userTickets)

	fmt.Println(("Enter your email: "))
	fmt.Scan(&email)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets, you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n",remainingTickets, conferenceName)

}
