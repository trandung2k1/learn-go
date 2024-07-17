package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring a single variable
	var age int                           // default = 0
	fmt.Println("My initial age is", age) // 0
	age = 23                              // assign value
	fmt.Println("My initial age is", age) // 23

	//Declaring a variable with an initial value
	var class string = "12B3"
	fmt.Println("My initial class is", class)

	//Type inference
	var name = "Tran Dung"
	fmt.Println("My initial name is", name)

	//Multiple variable declaration
	var price, quantity int = 5000, 10
	fmt.Println("Price is", price, "quantity is", quantity)

	// var price, quantity = 5000, 100

	var (
		fullName = "Tran Van Dung"
		height   = "1m75"
	)
	fmt.Println("Full name is", fullName, ",height is", height)

	//Short hand declaration
	count := 10
	fmt.Println("Count =", count)

	a, b := "Hello", 10 //short hand declaration
	fmt.Println("a is", a)
	fmt.Println("b is", b)

	c, d := 20, 30 // declare variables c and d
	fmt.Println("c is", c, "d is", d)
	d, e := 40, 50 // d is already declared but e is new
	fmt.Println("d is", d, "e is", e)

	min := math.Min(float64(c), float64(d))
	fmt.Println("Minimum value is", min)
}
