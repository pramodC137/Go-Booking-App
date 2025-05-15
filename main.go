package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const comferenceTickets = 50
	var remaniningTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", comferenceTickets, remaniningTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", comferenceTickets, remaniningTickets)
	fmt.Println("Get your tickets here toa attend")

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
