func maxProfit(prices []int) int {
    
    maxSoFar := 0
    max := 0
    for i := 1; i < len( prices); i++ {
        
        maxSoFar = maxInt( 0, maxSoFar + prices[ i] - prices[ i - 1])
        max = maxInt( max, maxSoFar)
    }
    return max
}

func maxInt( a int, b int) int {
    
    if a < b {
        return b
    }
    return a
}