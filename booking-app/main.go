package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	// Syntactical sugar in Golang
	// conferenceName := "Go Conference"
	// Here constants cannot be declared and explicit data types cannot be declared.

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)

	//fmt.Println("Welcome to", conferenceName, "booking application.")
	//fmt.Println("We have total of ", conferenceTickets, " tickets and remaining are ", remainingTickets)
	//fmt.Println("Please get your tickets here.")

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and remaining are %v.\n", conferenceTickets, remainingTickets)
	fmt.Println("Please get your tickets here.")

	var userName string
	var userTickets int
	// explicitly declaring variables. Needed when values are not assigned immediately

	userName = "Tom"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)
}
