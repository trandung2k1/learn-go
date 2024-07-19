package main

import (
	"fmt"
	"math"
)

func main() {
	// Declaring a constant
	const a = 50
	fmt.Println(a)

	//Declaring a group of constants
	const (
		retryLimit = 4
		httpMethod = "GET"
	)
	fmt.Println(retryLimit)
	fmt.Println(httpMethod)

	// a = 89 //reassignment not allowed

	var b = math.Sqrt(4) //allowed
	fmt.Println(b)

	// const c = math.Sqrt(4) //not allowed

	// String Constants, Typed and Untyped Constants
	const hello = "Hello World"
	fmt.Println(hello)
	const n = "Sam"
	var name = n
	fmt.Printf("type %T value %v", name, name)
	fmt.Println()
	var defaultName = "Sam" //allowed
	type myString string
	var customName myString = "Sam" //allowed
	// customName = defaultName        //not allowed
	customName = myString(defaultName) //allowed
	fmt.Println(customName)

	//Boolean Constants
	const trueConst = true
	type myBool bool
	var defaultBool = trueConst       //allowed
	var customBool myBool = trueConst //allowed
	// defaultBool = customBool          //not allowed
	defaultBool = bool(customBool)
	fmt.Println(defaultBool, customBool)

	// Numeric Constants
	const c = 5
	var intVar int = c
	var int32Var int32 = c
	var float64Var float64 = c
	var complex64Var complex64 = c
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// Numeric Expressions
	var d = 5.9 / 8
	fmt.Println(d)
}
