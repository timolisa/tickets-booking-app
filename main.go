package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var availableTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, availableTickets)

	// var bookings [50]string array has a fixed length in Go
	bookings := []string{}

	// fmt.Printf("conferenceName: %T, conferenceTickets: %T, availableTickets: %T\n", conferenceName, conferenceTickets, availableTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// a pointer is a variable that points to the memory location of another variable that references the actual value.
		// fmt.Println(availableTickets)
		// fmt.Println(&availableTickets)

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you wish to buy?")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 5 && len(lastName) >= 5
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(availableTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			availableTickets -= uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
			fmt.Printf("Number of available tickets: %v\n", availableTickets)
			fmt.Printf("Bookings: %v\n", bookings)

			// Blank identifiers are used to ignore a variable you don't want to use.
			// In Go, you have to make unused variables explicit.

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of the bookings: %v\n", firstNames)

			// var isTicketAvailable bool = availableTickets == 0

			if availableTickets == 0 {
				fmt.Println("Our tickets are sold out.")
				break
			}
		} else {
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", availableTickets, userTickets)
			if !isValidName {
				fmt.Println("First name or last name should be at least 5 characters long.")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email address")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter a valid ticket number")
			}
		}
	}
}

func greetUsers(conferenceName string, conferenceTickets int, availableTickets uint) {
	fmt.Println("Welcome to Event booking app")
	fmt.Printf("Welcome to %v conference booking app\n", conferenceName)
	fmt.Print("total tickets: ", conferenceTickets, "available tickets: ", availableTickets)
	fmt.Println("Get your tickets here to attend.")
}