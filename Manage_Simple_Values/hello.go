package main

//import "fmt" does the same if you have only one library to import
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	usingReaders()
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
