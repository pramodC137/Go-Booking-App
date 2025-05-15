package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const comferenceTickets = 50
	var remaniningTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", comferenceTickets, remaniningTickets)
	fmt.Println("Get your tickets here toa attend")

	fmt.Println(conferenceName)

}
