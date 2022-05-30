func rob(nums []int) int {
    
    if len( nums) == 1 {
        
        return nums[ 0]
    }
    
    solution := make( []int, len( nums), len(nums))
    solution[ len( solution) - 1] = nums[ len( solution) - 1]
    solution[ len( solution) - 2] = maxInt( nums[ len( solution) - 1], nums[ len( solution) - 2] )
    
    for i := len( solution) - 3; i >= 0; i-- {
        
        curSum := solution[ i + 2] + nums[ i]
        if curSum > solution[ i + 1] {
            
            solution[ i] = curSum
            
        } else {
            
            solution[ i] = solution[ i + 1]
        }
        
    }
    return solution[ 0]
}

func maxInt( a int, b int) int {
    
    if a > b {
        
        return a
    }
    return b
}