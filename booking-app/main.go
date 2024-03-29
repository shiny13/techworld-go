package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

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

	//var bookings [50]string //declared as an array
	var bookings []string //declared as a slice
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// explicitly declaring variables. Needed when values are not assigned immediately

	// take a user input using the fmt package
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email:")
	fmt.Scan(&email)
	fmt.Println("Enter the amount of tickets:")
	fmt.Scan(&userTickets)

	//bookings[0] = firstName + " " + lastName //adding values to an array
	bookings = append(bookings, firstName+" "+lastName)

	/*fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))*/

	fmt.Printf("Thank you %v %v for booking %v tickets.\nYou will receive a confirmation email at %v soon.\n", firstName, lastName, userTickets, email)
	fmt.Printf("All the bookings: %v\n", bookings)
}
