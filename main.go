package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const comferenceTickets = 50
	var remaniningTickets uint = 50
	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", comferenceTickets, remaniningTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", comferenceTickets, remaniningTickets)
	fmt.Println("Get your tickets here toa attend")

	for remaniningTickets > 0 && len(bookings) < 50 {
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

		if userTickets < remaniningTickets {

			remaniningTickets = remaniningTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			//arrsy details
			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array type: %T\n", bookings)
			fmt.Printf("Array lenght: %v\n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaning for %v\n", remaniningTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("First name of bookings are %v\n", firstNames)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			if remaniningTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}

		} else if userTickets == remaniningTickets {

		} else {
			fmt.Printf("We only have %v tickets remaning, so you can't book %v tickets\n", remaniningTickets, userTickets)
			continue
		}
	}
}
