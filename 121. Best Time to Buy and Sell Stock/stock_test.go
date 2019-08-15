package stock_test

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	min := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price > min {
			profit := price - min
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			min = price
		}
	}

	return maxProfit
}
