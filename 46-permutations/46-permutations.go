func permute(nums []int) [][]int {
    
    used := make( []bool, len( nums), len(nums))
    curSolution := make( []int, len( nums), len(nums))
    solutions := make( [][]int, 0)
    
    return calculate( len( nums), nums, used, curSolution, solutions)
}

func calculate( flag int, nums []int, used []bool, curSolution []int, solutions [][]int) [][]int {
    
    if flag == 0 {
        
        solution := make( []int, len( nums))
        copy( solution, curSolution)
        return append( solutions, solution)
    } else {
        
        for i := range nums {
            
            if !used[ i] {
                
                used[ i] = true
                curSolution[ flag - 1] = nums[ i]
                solutions = calculate( flag - 1, nums, used, curSolution, solutions)
                used[ i] = false
            }
        }
    }
    return solutions
}