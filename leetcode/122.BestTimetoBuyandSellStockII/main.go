package main

import "math"

//func maxProfit(prices []int) int {
//	result := 0
//	for i, price := range prices {
//		for j := i; j < len(prices); j++ {
//			if prices[j] > price {
//				result = utils.Max(result, prices[j]-price)
//			}
//		}
//	}
//	return result
//}

func calcProfit(prices []int, idx int) int {
	smallestPrice := math.MaxInt32
	result := 0

	for i := idx; i < len(prices); i++ {
		if prices[i] < smallestPrice {
			smallestPrice = prices[i]
		}
		profit := prices[i] - smallestPrice
		if profit > result {
			result = profit
		}
	}
	return result
}

func maxProfit(prices []int) int {
	length := len(prices) - 1
	i := 0
	profit := 0
	for i < length {
		for i < length && prices[i+1] < prices[i] {
			i++
		}
		buy := prices[i]

		for i < length && prices[i+1] > prices[i] {
			i++
		}
		sell := prices[i]
		profit += sell - buy
	}

	return profit
}
