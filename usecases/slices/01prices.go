package slices

import "fmt"

//** Problem:
// Given a slice of prices of items in a store, you need a function to calculate the total purchase amount.
//
//** Description:
// You're provided with a slice containing prices of items purchased by a customer in a store. You're required
// to implement a function that takes this slice of prices and returns the total amount the customer owes for all the purchased items.
//
//** Input:
// A slice of item prices in the store. For example:
// prices := []int{50, 30, 20, 15, 25}
//
//** Output:
// The total purchase amount, which is the sum of all the prices in the slice.
//
//** Example:
// If you have the price slice prices := []int{50, 30, 20, 15, 25}, the function should return 140,
// as it's the sum of all elements in the slice (50 + 30 + 20 + 15 + 25 = 140).
//

func Prices() {
	prices := []int{50, 30, 20, 15, 25}
	total := CalculateTotalPurchase(prices)
	fmt.Println(total)
}

func CalculateTotalPurchase(prices []int) int {
	total := 0
	for _, v := range prices {
		total += v
	}
	return total
}
