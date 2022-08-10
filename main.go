package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

func main() {
	conferenceName := "Go Conference" // applies to only variables
	const conferenceTickets = 50
	var remainingTickets = 50
	// bookings := make([]map[string]string, 0)
	type UserData struct {
		firstName              string
		lastName               string
		email                  string
		numberOfTickets        uint
		isOptedInForNewsLetter bool
	}
	bookings := make([]UserData, 0)

	greetUsers(conferenceName)

	fmt.Printf("Conference name is %T, remainingTickets is %T\n", conferenceName, remainingTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are availaible")
	fmt.Println("Get Your tickets here")

	for {
		firstName, lastName, email, userTickets := takeInput()
		// var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		// isValidEmail := strings.Contains(email, "@")

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNumber {
			fmt.Printf("We have only %v tickets remaining\n", userTickets)
			continue
		}

		if isValidName && isValidEmail {

			remainingTickets = remainingTickets - (userTickets)

			// fmt.Printf("The whole slice: %v\n",bookings)
			// fmt.Printf("The first value %v\n",bookings[0])
			// fmt.Printf("Slice type %T\n",bookings)
			// fmt.Printf("Slice type %v\n",len(bookings))

			firstNames := []string{}

			// var userData = make(map[string]string)
			// userData["firstName"] = firstName
			// userData["lastName"] = lastName
			// userData["email"] = email
			// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: uint(userTickets),
			}
			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)

			go sendTicket(uint(userTickets), firstName, lastName, email)

			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName)
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve tickets at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			if remainingTickets == 0 {
				fmt.Println("Our conference has no tickets.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email format is invalid")
			}

		}

	}
}

func greetUsers(conferenceName string) {
	fmt.Println("Welcome to", conferenceName)
}

func takeInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###################")
}
