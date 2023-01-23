package main


import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
    // fmt.Println(conferenceName)

    var remainingTickets uint = 50
	
	bookings := []string{}
    
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)


	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")

    fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets) 

	fmt.Println("Get to our tickets to get here!")

    for remainingTickets > 0 && len(bookings) < 50 {
		
		
		var firstName string
		var lastName  string
		var email  string
		var userTickets uint
		
		// ask user for their name
		fmt.Println("Enter your firstName:") 	
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your lastName:")
		fmt.Scan(&lastName)
		
		
		fmt.Println("Enter no of tickets:")
		fmt.Scan(&userTickets)
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		
		// fmt.Println(remainingTickets)
		// fmt.Println(&remainingTickets)
        
		isValidName :=  len(firstName) >= 2 && len(lastName) >=2  
        isValidEmail := strings.Contains(email,"@")
        isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets  

    
		

        if isValidName && isValidEmail && isValidTicketNumber  {
		    remainingTickets = remainingTickets - userTickets
		
		// bookings[0] = firstName + " " + lastName 
		bookings = append(bookings, firstName + "" + lastName)
		
		// userName = "Tom"
		// userTickets = 2
		
		fmt.Printf("The whole slice:%v\n",bookings) 
		fmt.Printf("The first slice:%v\n",bookings[0])
		fmt.Printf("Slice types:%T\n",bookings)
		fmt.Printf("Slice length: %v\n",len(bookings))
		
		
		fmt.Printf("Thank you %v %v for booking %v the tickets, You will receiving the email    %v tickets \n", firstName,lastName,userTickets,email)	
		
		fmt.Printf("%v tickets remaining %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
        for _, booking := range bookings {
			var names = strings.Fields(booking)
			
			firstNames = append(firstNames, names[0]) 
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	

        

	    if remainingTickets == 0 {
			// end program
    
			fmt.Println("Our conference is book out.Come back next year.")
            break
		}	


		}   else {

			fmt.Println("You input data is invalid and try again")
			
			
		}
		
	}
	
		
}
