package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName string = "Go conference"

const conferenceTickets int = 50

var availableTickets uint = 50

var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, availableTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(firstName, lastName, email, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings: %v\n", firstNames)

			if availableTickets == 0 {
				fmt.Println("Our tickets are sold out.")
				break
			}
		} else {
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

func greetUsers() {
	fmt.Println("Welcome to Event booking app")
	fmt.Printf("Welcome to %v conference booking app\n", conferenceName)
	fmt.Print("total tickets: ", conferenceTickets, "available tickets: ", availableTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you wish to buy?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func bookTickets(firstName string, lastName string, email string, userTickets int) {
	availableTickets -= uint(userTickets)

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Number of available tickets: %v\n", availableTickets)
	fmt.Printf("Bookings: %v\n", bookings)
}