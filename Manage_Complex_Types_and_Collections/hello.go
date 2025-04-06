package main

import (
	"fmt"
	"sort"
)

func main() {
	// new() only creates a pointer to the memory but does not allocate it and does not create a default value
	// make() allocates and initialize memory and create a default value
	/*
		For example this is not a good code:
		var m map[string]int
		m["key"] = 42
		fmt.Println(m)

		in this case the code will panic when we want ot assign a value because we did not allocate memory
	*/
	m := make(map[string]int)
	m["Key"] = 42
	fmt.Println(m)

	anInt := 42
	var p *int = &anInt
	if p == nil {
		fmt.Println("P is nil")
	} else {
		fmt.Println("The value of p: ", *p)
	}

	// We can see that pointers work like in any other language
	value1 := 42.13
	pointer1 := &value1
	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer: ", *pointer1)
	fmt.Println("Original value: ", value1)

	// Arrays
	arrays()

	//Slices
	slices()

	// Maps
	maps()

	// using structs
	poodle := Dog{"Poodle", 12}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle) // this can write out the name of the field what was addressed
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Weight = 29
	fmt.Println(poodle)
}

func arrays() {
	// Arrays are works like basic Cpp lists size is constantly allocated, and you have pretty small room to wiggle
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	fmt.Println("Number of colors: ", len(colors))
	fmt.Println("Number of numbers: ", len(numbers))
}

func slices() {
	// if you don't put a number in the [] than it is a slice, works like std::vector
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// But better to allocate memory for them like this for efficiency:
	var colors2 = make([]string, 0, 3)
	colors2 = append(colors2, "Red", "Green", "Blue")
	// in this case the slice will be re-allocated
	colors2 = append(colors2, "Purple", "Pink")
	fmt.Println(colors2)

	// remove element
	colors2 = removeFromSlice(colors2, 0)
	fmt.Println(colors2)

	// Sorting
	sort.Strings(colors2)
	fmt.Println(colors2)
}

func removeFromSlice(slice []string, i int) []string {
	// ... operator is the same as in js or rust
	// [:i] means that we get a slice with the first n elements i not included
	// slice[i+1]... means we get the slice's elements after i+1, i+1 is included in this case
	return append(slice[:i], slice[i+1:]...)
}

func maps() {
	// Maps are UnorderedMaps in go, when you print them it should be alphabetical but you should ont depend on that it won't be like that every time
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "NewYork"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// Make sure we print them alphabetically
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}

type Dog struct {
	Breed  string
	Weight int
}
