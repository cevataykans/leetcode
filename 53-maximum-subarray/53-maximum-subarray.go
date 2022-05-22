func maxSubArray(nums []int) int {
    
    maxSoFar, maxEndingHere := nums[ 0], nums[ 0]
    for i := 1; i < len( nums); i++ {
        
        maxEndingHere = max( maxEndingHere + nums[ i], nums[ i])
        maxSoFar = max( maxSoFar, maxEndingHere)
    }
    return maxSoFar
}

func max( f, s int) int {
    
    if f > s {
        return f
    }
    return s
}