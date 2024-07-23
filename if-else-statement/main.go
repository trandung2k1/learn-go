package main

import "fmt"

func main() {

	// If statement syntax
	// num := 10
	// if num%2 == 0 {
	// 	fmt.Println("The number", num, "is even")
	// 	return
	// }
	// fmt.Println("The number", num, "is odd")

	//If else statement
	num := 10
	if num%2 == 0 {
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}

	// If … else if … else statement
	// age := 10
	// ticketPrice := 0
	// if age < 5 {
	// 	ticketPrice = 0
	// } else if age >= 5 && age <= 22 {
	// 	ticketPrice = 10
	// } else {
	// 	ticketPrice = 15
	// }
	// fmt.Printf("Ticket price is $%d", ticketPrice)

	// If with assignment
	ticketPrice := 0
	if age := 10; age < 5 {
		ticketPrice = 0
	} else if age >= 5 && age <= 22 {
		ticketPrice = 10
	} else {
		ticketPrice = 15
	}
	fmt.Printf("Ticket price is $%d", ticketPrice)

}
