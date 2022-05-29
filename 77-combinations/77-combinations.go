func combine(n int, k int) [][]int {
    
    curSolution := make( []int, 0)
    solutions := make( [][]int, 0)
    
    return findCombinations( k, 1, n, curSolution, solutions)
}

func findCombinations( flag int, cur int, n int, curSolution []int, solutions [][]int) [][]int {
    
    if flag == 0 {
        
        solution := make( []int, len( curSolution))
        copy( solution, curSolution)
        return append( solutions, solution)
        
    } else {
        
        for i := cur; i <= n; i++ {
            
                curSolution = append( curSolution, i)
                solutions = findCombinations( flag - 1, i + 1, n, curSolution, solutions)
                curSolution = curSolution[ :len(curSolution) - 1]
        }
    }
    return solutions 
}