package main

import "fmt"

func main() {
	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	var bananas = CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Println(result)
}

func calculateTotal(cart []CartItem) float64 {
	var value float64 = 0
	for _, cart := range cart {
		value += cart.Price * float64(cart.Quantity)
	}
	return value
}

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}
