package main

import (
	"fmt"
	"time"
)

func main() {
	//theAnswer := 0
	var result string

	// The theAnswer variable will only be usable in the if statement's context
	if theAnswer := 42; theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day number is %v\n", dayNumber)

	// In switches there is no need for break it is like Java
	switch dayNumber = 1; dayNumber {
	case 1:
		result = "It's a Monday!"
	case 2:
		result = "It's a Tuesday!"
	case 3:
		result = "It's a Wednesday!"
	case 4:
		result = "It's a Thursday!"
	case 5:
		result = "It's a Friday!"
	default:
		result = "It's Weekend!"
	}

	fmt.Println(result)

	x := -42
	switch {
	case x < 0:
		result = "Less than 0"
	case x == 0:
		result = "Equals 0"
	case x > 0:
		result = "Bugger than 0"
	}

	fmt.Println(result)

	// we can fallthrough the cases if needed
	switch result = ""; {
	case x < 0:
		result = result + "This "
		fallthrough
	case x == 0:
		result = result + "is "
		fallthrough
	case x > 0:
		result = result + "fallthrough!"
	}

	fmt.Println(result)

	colors := []string{"Red", "Green", "Blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	// the fist one is the index and the second is the value if it isn't a map
	for _, color := range colors {
		fmt.Println(color)
	}

	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	for state, _ := range states {
		fmt.Println(states[state])
	}

	// this is a while loop in go
	value := 0
	sum := 0
	for value < 5 {
		sum += value
		fmt.Printf("The value is %v\n", value)
		fmt.Printf("The sum is %v\n", sum)
		value++
	}
}
