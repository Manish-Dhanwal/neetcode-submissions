func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0


	for _, v := range prices{
		if v < minPrice {
			minPrice = v
		}
		profit := v - minPrice
		if profit > maxProfit{
			maxProfit = profit
		} 
	} 
	return maxProfit
}
