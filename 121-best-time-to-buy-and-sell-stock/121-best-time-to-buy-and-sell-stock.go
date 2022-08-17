func maxProfit(prices []int) int {
    
    profit := 0
    buyPrice := prices[ 0]
    for i := 1; i < len(prices); i++ {
        if prices[ i] < buyPrice {
            buyPrice = prices[ i]
        } else if profit < prices[ i] - buyPrice {
            profit = prices[ i] - buyPrice
        }
    }
    return profit
}