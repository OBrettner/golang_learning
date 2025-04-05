package main

//import "fmt" does the same if you have only one library to import
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	// if it is undeclared it will be default
	var anInteger int = 23
	var aString string = "hello"

	fmt.Println(aString)
	fmt.Println(anInteger)

	//implicit declaration, not sure why var is not needed anymore
	aDouble := 32.3

	fmt.Println(aDouble)

	//this is the same const as in cpp
	const constInteger int = 42

	// =========TYPES========
	// string
	// bool

	// ========NUMERICS=======
	// uint8
	// uint16
	// uint32
	// uint64
	// int8
	// int16
	// int32
	// int64

	// byte
	// uint <- 32/64
	// int <- 32/64
	// uintptr

	// float32
	// float64

	// complex64
	// complex128

	//=========COMPLEX TYPES ========
	// Array
	// Slice
	// Map
	// Struct

	// Function
	// Interface
	// Channel

	// Pointer
	printExample()

	// Input from console
	//usingReaders()

	// Math operations
	mathOperations()

	// Dates and times
	datesAndTimes()
}

func printExample() {
	// Example for the fmt Print function
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy dog"
	number1 := 32

	fmt.Println(str1, str2, str3)
	strLength, err := fmt.Println("The number is", number1)
	if err == nil {
		fmt.Println("String length:", strLength)
	}

	fmt.Printf("Value of number: %v\n", number1)

	fmt.Printf("Data type: %T\n", number1)
}

func usingReaders() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
}

func mathOperations() {
	// There is no implicit conversion nin go
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Iteger sum: ", intSum)

	f1, f2, f3 := 23.5, 45.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	// For conversion, we have functions
	total := float64(i1) * f2
	fmt.Println("total: ", total)

	// Let's start using the Math package
	floatSum = math.Round(floatSum*100) / 100
	fmt.Printf("Float sum: %v\n\n", floatSum)

	// the PI constant
	fmt.Println("The value of PI is: ", math.Pi)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumference: %.2f\n", circumference)
}

func datesAndTimes() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	n := time.Now()
	fmt.Printf("The time currently is %s\n", n)
	fmt.Printf("This object type is %T\n", n)

	fmt.Printf(n.Format(time.ANSIC) + "\n")

	tomorrow := n.AddDate(0, 0, 1)
	fmt.Printf(tomorrow.Format(time.ANSIC) + "\n")

	format := "Mon 2006-02-01"
	fmt.Printf(tomorrow.Format(format) + "\n")
}
