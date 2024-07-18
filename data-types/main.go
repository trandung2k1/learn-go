package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Default value = false
	var isValid bool
	fmt.Println(isValid)

	// int8 - 8bits, -128 -> 127, (2^7)-1
	var number int8 = 127
	fmt.Println(number)

	// int16 - 16bits, (2^15)-1
	// int32 - 32bits, (2^31)-1
	// int64 - 64bits, (2^63)-1
	// int : int32 or int64

	fmt.Printf("type of a is %T, size of a is %d bytes", number, unsafe.Sizeof(number))

	// uint8 - 8bits, 0 -> 255, (2^8)-1
	// uint16 - 16bits, (2^16)-1
	// uint32 - 32bits, (2^32)-1
	// uint64 - 64bits, (2^64)-1
	// uint : uint32 or uint64

	// float32 - 32bits,
	// float32 - 64bits,

	fmt.Println()
	var mark float32 = 5.1
	//%g
	fmt.Printf("%f\n", mark)

	// complex
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	// string
	first := "Tran"
	last := "Dung"
	name := first + " " + last
	fmt.Println("My name is", name)

	//type conversion
	a := 80   //int
	b := 91.8 //float64
	sum := a + int(b)
	fmt.Println(sum)

	i := 10
	var j float64 = float64(i)
	fmt.Println("j =", j)

	// byte
	var a1 byte = 97
	fmt.Printf("%c\n", a1) //a

	var a2 []byte
	a2 = append(a2, 98)
	a2 = append(a2, 99)
	fmt.Printf("%c", a2) //[b c]

	s := "hello"
	r := []rune(s)
	fmt.Println(r)                         //[104 101 108 108 111]
	fmt.Println(utf8.RuneCountInString(s)) //5
	fmt.Println(string(r))                 //hello
}
