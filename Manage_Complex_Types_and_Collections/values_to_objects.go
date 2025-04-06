package main

import "fmt"

func main() {
	colorNames := []string{"Red", "Green", "Blue"}
	hexValues := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colors := slicesToObjects(colorNames, hexValues)
	fmt.Println(colors)
}

func slicesToObjects(colorNames []string, hexValues []int) []Color {
	if len(colorNames) != len(hexValues) {
		panic("Not equal length")
	}

	colorSlice := make([]Color, 0, len(colorNames))
	for i := range colorNames {
		colorSlice = append(colorSlice, Color{Name: colorNames[i], Hex: hexValues[i]})
	}
	return colorSlice
}

type Color struct {
	Name string
	Hex  int
}
