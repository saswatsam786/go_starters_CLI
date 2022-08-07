package main

import (
	"fmt"
	"strings"
)

func main()  {
	conferenceName :="Go Conference" // applies to only variables
	const conferenceTickets = 50
	var remainingTickets uint = 50
	 bookings:= []string{}

	fmt.Printf("Conference name is %T, remainingTickets is %T\n",conferenceName,remainingTickets)

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Println("We have total of",conferenceTickets,"tickets and",remainingTickets,"are availaible")
	fmt.Println("Get Your tickets here")

	for{
		var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets-(userTickets)
	bookings=append(bookings,firstName+" "+lastName)

	// fmt.Printf("The whole slice: %v\n",bookings)
	// fmt.Printf("The first value %v\n",bookings[0])
	// fmt.Printf("Slice type %T\n",bookings)
	// fmt.Printf("Slice type %v\n",len(bookings))

		firstNames:=[]string{}

		for _,booking := range bookings {
			var names = strings.Fields(booking)
			firstNames=append(firstNames,names[0])
		}


	fmt.Printf("The first names of bookings are: %v\n",firstNames)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve tickets at %v \n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

	}

	
}