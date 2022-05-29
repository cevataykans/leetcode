func isHappy(n int) bool {
 
    visited := make( map[int]bool)
    tempNum := n
    for {
        
        tempNum = getSquares( tempNum)
        if tempNum == 1 {
            return true
        }
        
        if visited[ tempNum] {
            break
        }
        visited[ tempNum] = true
    }
    return false
}

func getSquares( num int) int {
    
    curSum := 0
    
    for num != 0 {
        
        digit := num % 10
        num = num / 10
        curSum += digit * digit
    }
    return curSum
}