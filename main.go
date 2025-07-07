package main

import (
	"booking-app/helper"
	"fmt"
)

const comferenceTickets = 50

var conferenceName = "Go Conference"
var RemaniningTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", comferenceTickets, RemaniningTickets, conferenceName)

	greetUers()

	for RemaniningTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isvalidTicketNumber := helper.ValidateUserInputs(firstName, lastName, email, userTickets, RemaniningTickets)

		if isValidName && isValidEmail && isvalidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames()
			fmt.Printf("First name of bookings are %v\n", firstNames)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			if RemaniningTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else if userTickets == RemaniningTickets {

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isvalidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func greetUers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", comferenceTickets, RemaniningTickets)
	fmt.Println("Get your tickets here toa attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their details
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	RemaniningTickets = RemaniningTickets - userTickets

	// create a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	//arrsy details
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array lenght: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaning for %v\n", RemaniningTickets, conferenceName)
}
