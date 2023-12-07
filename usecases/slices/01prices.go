package slices

import (
	"fmt"
	"runtime"
	"sync"
)

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
	numOfCPU := runtime.NumCPU()
	prices := []int{50, 30, 20, 15, 25, 10, 10, 10, 10}
	total := 0
	if len(prices) > numOfCPU {
		total = CalculateTotalPurchaseRoutines(prices, numOfCPU)
	} else {
		total = CalculateTotalPurchase(prices)
	}
	fmt.Println(total)
}

func CalculateTotalPurchase(prices []int) int {
	total := 0
	for _, v := range prices {
		total += v
	}
	return total
}

func CalculateTotalPurchaseRoutines(prices []int, numWorkers int) int {

	priceChunks := divideByPrices(prices, numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	totalChan := make(chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(chunk []int) {
			defer wg.Done()
			partialTotal := 0
			for _, v := range chunk {
				partialTotal += v
			}
			totalChan <- partialTotal
		}(priceChunks[i])
	}

	go func() {
		wg.Wait()
		close(totalChan)
	}()

	total := 0
	for partial := range totalChan {
		total += partial
	}

	return total
}

func divideByPrices(prices []int, numWorkers int) [][]int {
	result := make([][]int, numWorkers)
	chunkSize := len(prices) / numWorkers
	extra := len(prices) % numWorkers

	startIndex := 0
	for i := 0; i < numWorkers; i++ {
		var currentChunkSize int
		if i < extra {
			currentChunkSize = chunkSize + 1
		} else {
			currentChunkSize = chunkSize
		}

		endIndex := startIndex + currentChunkSize
		if endIndex > len(prices) {
			endIndex = len(prices)
		}

		result[i] = prices[startIndex:endIndex]
		startIndex = endIndex
	}

	return result
}

func hasDecimals(number float64) bool {
	return number != float64(int(number))
}
