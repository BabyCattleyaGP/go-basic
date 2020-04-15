package main

import "fmt"

func main() {
	const earthsGravity = 9.80665

	var number int8
	number = 30

	var floatt float32 = 5.8

	var tryString string = "Baso"

	days := 6
	var hours = 24

	//all string use +, otherwise can use ,
	fmt.Println("Day spent in hours", days*hours)

	fmt.Println("My favorite snack is " + tryString)
	fmt.Println("Print number from var:", number)
	fmt.Println("Print float number from var:", floatt)
	fmt.Println("The earth gravity is", earthsGravity)
	fmt.Println("23 multiplies by 4 is", 23*4)
}
