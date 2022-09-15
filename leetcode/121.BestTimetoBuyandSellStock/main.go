package _21_BestTimetoBuyandSellStock

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

func maxProfit(prices []int) int {
	smallestPrice := math.MaxInt32
	result := 0

	for _, price := range prices {
		if price < smallestPrice {
			smallestPrice = price
		}
		profit := price - smallestPrice
		if profit > result {
			result = profit
		}
	}
	return result
}
