package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
	println()

	// Break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	println()

	// Continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	println()

	//infinite loop

	// for {
	// 	fmt.Println("Hello World")
	// }

	// Nested for loops
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
