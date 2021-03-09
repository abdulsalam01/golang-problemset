package main

import "fmt"

func MaxProfit(prices []int, i int) int {
	// sorting first
	for i := 1; i < len(prices); i++ {
		for j := i; j > 0; j-- {
			if prices[j-1] > prices[j] {
				tmp := prices[j-1]
				prices[j-1] = prices[j]
				prices[j] = tmp
			}
		}
	}

	// sum all
	max := sumArray(prices[len(prices)-i:])
	min := sumArray(prices[0:i])

	return (max - min)
}

// helper
func sumArray(arr []int) int {
	res := 0

	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}

	return res
}

func main() {
	// test case - optimum and minimum element finder
	fmt.Print(MaxProfit([]int{4, 11, 2, 20, 59, 80}, 2))
}
