package main

import "fmt"

func main() {
	var rs = sum(1, 2)
	fmt.Println("Sum:", rs)
	area, perimeter := rectProps(10.8, 5.6)
	area1, perimeter1 := rectProps1(10.8, 5.6)
	area2, _ := rectProps1(10.8, 5.6)

	fmt.Printf("Area %f Perimeter %f", area, perimeter)
	fmt.Println()
	fmt.Printf("Area1 %f Perimeter1 %f", area1, perimeter1)
	fmt.Println()
	fmt.Printf("Area2 %f", area2)
}

func sum(a, b int) int {
	return a + b
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps1(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}
