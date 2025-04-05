package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value1 := "10"
	value2 := "5.5"
	result := calculate(value1, value2)
	fmt.Println("Result:", result)

}

func calculate(value1 string, value2 string) float64 {
	f1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err1 != nil {
		panic(err1)
	}
	f2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err2 != nil {
		panic(err2)
	}
	return f1 + f2
}
