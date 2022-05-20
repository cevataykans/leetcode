func sortedSquares(nums []int) []int {
    
    squared := make( []int, len( nums))
    start := 0
    end := len( nums) - 1
    curI := end
    for start <= end {
        
        if square( nums[ start] ) > square( nums[ end] ) {
            
            squared[ curI] = square( nums[ start] )
            start++
            
        } else {
            
            squared[ curI] = square( nums[ end])
            end--
        }
        curI--
    }
    return squared
}

func square( num int) int {
    return num * num
}