package main

import "fmt"

func main() {
	fmt.Println("welcome to conference booking system!!!!")

	var tickets int = 100
	var ticketList = make(map[string]int)
start:
	fmt.Println("Choose from given options :")
	fmt.Println("1) Book tickets")
	fmt.Println("2) Show booking list")
	fmt.Println("3) exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {

	case 1:
	take_count:
		var userName string
		fmt.Println("Enter your Name")
		fmt.Scan(&userName)
		fmt.Println("------------------------------------------------------------------------------")
		fmt.Println("welcome,", userName)
		fmt.Println("Available tickets :", tickets)
		fmt.Print("Enter number of tickets you want to book :")
		var number_of_tickets int
		fmt.Scan(&number_of_tickets) 
		
		if (number_of_tickets > 0 && number_of_tickets < tickets) {
			ticketList[userName]=number_of_tickets
			tickets -= number_of_tickets
		} else if (number_of_tickets > tickets) {
			fmt.Println("Please Check available number of tickets and enter the value.")
			goto take_count
		} else {
			fmt.Println("Enter a valid number!!!!")
			goto take_count
		}
		goto start

	case 2:
		fmt.Println("Bookng List:")
		for k , v := range ticketList {
			fmt.Printf("%v \t %v",k,v)
		}
		goto start

	case 3:
		fmt.Println("Thank You..........")
	}

}
