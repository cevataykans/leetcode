func minimumTotal(triangle [][]int) int {
    
    solution := make( []int, len( triangle))
    copy( solution, triangle[ len(triangle) - 1])
    
    for i := len( triangle) - 2; i >= 0; i-- {
        
        for j := 0; j <= i; j++ {
            
            solution[ j] = minInt( solution[ j], solution[ j + 1]) + triangle[ i][ j]
        }
    }
    return solution[ 0]
}

func minInt( a int, b int) int {
    
    if a < b {
        
        return a
    }
    return b
}