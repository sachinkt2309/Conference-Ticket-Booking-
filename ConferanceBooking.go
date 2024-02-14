package main

import (
	//	"bufio"
	"booking-app/validate"
	"fmt"
	"sync"
	"time"
	//	"log"
	//	"os"
)

var pl = fmt.Println
var pf = fmt.Printf

var conferanceName string = "Go Coding"

const conferanceTickets int = 50

var remainingTickets uint = 50
var booking = make([]UserData, 0)

type UserData struct {
	firstName   string // defining datatype string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greeting()

	//printf example
	//print datatypes
	// pf("conferance Tickets is %T, remaining Tickets is %T, conferanceName is %T \n", conferanceTickets, remainingTickets, conferanceName)

	//infinite For loop
	//for remainingTickets > 0 && len(booking) <= 50 {

	// function to get user data
	firstName, lastName, email, userTickets := getUserInput()

	// function to validate user data
	isValidName, isValidEmail, isValidTickets := validate.ValidateUser(firstName, lastName, email, userTickets, remainingTickets)

	if !isValidName {
		pl("Please Enter Vaild Name\n")
	} else if !isValidEmail {
		pl("Please Enter Valid Email\n")

	} else if isValidTickets {

		bookTickets(firstName, lastName, email, userTickets)

		// created new thread to run sendingTicket method
		wg.Add(1) // synchronization call
		go sendingTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstName()
		pf(" booking users first names are: %v \n\n", firstNames)

		if remainingTickets == 0 {
			pl("all tickets are sold out please come back next time")
			//break
		}
	} else {
		pl("Enter ticket number is not Valid please try again\n")
	}

	//}
	wg.Wait()
}
func greeting() {
	pf("Welcome to our %v Booking Application\n", conferanceName)

	pf("we have total of %v tickets and %v are still available\n", conferanceTickets, remainingTickets)

}
func getFirstName() []string {
	firstNames := []string{}

	//for each loop or Range in loop
	for _, book := range booking {
		//var names = strings.Fields(user)
		firstNames = append(firstNames, book.firstName)
	}
	return firstNames

}

func getUserInput() (string, string, string, uint) {
	var firstName string // defining datatype string
	var lastName string
	var email string
	var userTickets uint // defined int type
	// user input
	pl("Enter your First Name:")
	fmt.Scan(&firstName)
	pl("Enter your Last Name")
	fmt.Scan(&lastName)
	pl("Enter your email address:")
	fmt.Scan(&email)
	pl("Enter no of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {

	// Map implementation

	/* var userData = make(map[string]string)
	userData["FirstName"] = firstName
	userData["LastName"] = lastName
	userData["email"] = email
	userData["NoOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)   */

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	booking = append(booking, userData)

	pf("List of booking is: %v\n", booking)

	remainingTickets = remainingTickets - userTickets

	pf("thank you %v for booking %v tickets, you will recive confirmation at %v email \n\n ", firstName, userTickets, email)

	pl("out of: ", conferanceTickets, " remaining tickets are: ", remainingTickets)
}
func sendingTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	pl("-------------------------")
	pf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
	pl("-------------------------")
	wg.Done()
}
