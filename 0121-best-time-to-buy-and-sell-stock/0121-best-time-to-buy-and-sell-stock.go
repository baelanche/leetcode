func maxProfit(prices []int) int {
    max := 0
    buy := prices[0]

    for _, price := range prices {
        if price < buy {
            buy = price
        } else {
            profit := price - buy
            if max < profit {
                max = profit
            }
        }
    }

    return max
}