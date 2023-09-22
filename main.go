package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// package level variables (cant't use := symbol)
var conferenceName string = "Go Conference"

const numberOfTickets uint = 100

var remainingTickets uint = numberOfTickets

var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	for {
		greetuser()

		firstName, lastName, email, userTickets := userInput()
		isValidEmail, isValidName, isValidTicket := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicket {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			if remainingTickets == 0 {
				fmt.Println("All tickets are booked, thank you!")
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("Enter a valid email id")
			}
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidTicket {
				fmt.Printf("Less than %v tickets are available\n", userTickets)
			}
			continue
		}
	}
	wg.Wait()
}

func greetuser() {
	fmt.Println("            Welcome to", conferenceName)
	fmt.Printf("----------------#################------------------\n\n")
	fmt.Printf("Get your tickets to attend ----> Available tickets: %v\n\n", remainingTickets)
}

// function to return multiple values
func userInput() (string, string, string, uint) {
	var firstName string
	fmt.Print("Please provide your first name: ")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Print("Please provide your last name: ")
	fmt.Scan(&lastName)

	var email string
	fmt.Print("Please provide your email: ")
	fmt.Scan(&email)

	var userTickets uint
	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func printFirstName() {
	firstNames := []string{}
	for _, booking := range bookings {
		name := strings.Fields(booking.firstName)
		firstNames = append(firstNames, name[0])
	}
	fmt.Println("First name of the bookings:", firstNames)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Println(bookings)
	fmt.Printf("Thank you %v for booking %v tickets, you receive a confirmation email at %v\n\n", firstName, userTickets, email)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################################")
	fmt.Printf("Sending ticket:\n%v \nto email %v\n", ticket, email)
	fmt.Println("#################################")
	wg.Done()
}
