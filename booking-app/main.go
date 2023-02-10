package main

import "fmt"

func main() {

	// Defining a variable
	var conferenceName = "Go conference"
	// Constants do not change
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Println("Welcome to the", conferenceName, "booking application\n Buy your tickets here.")
	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v and are %v available.", conferenceTickets, remainingTickets)

}
