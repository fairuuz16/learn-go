package main

import (
	"fmt"
	// "strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := printFirstNames()
		fmt.Printf("The first names %v\n", firstNames)

		if remainingTickets == 0 {
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

// func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidEmail := strings.Contains(email, "@")
// 	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
// 	return isValidName, isValidEmail, isValidTicketNumber
// }

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
