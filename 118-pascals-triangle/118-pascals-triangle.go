func generate(numRows int) [][]int {
    
    solution := [][]int{ []int{ 1 } }
    if numRows == 1 {
        
        return solution
    }
    
    solution = append( solution, []int{ 1, 1 } )
    if numRows == 2 {
        
        return solution
    }
    
    for i := 3; i <= numRows; i++ {
        
        newRow := []int{ 1 }
        lastRow := solution[ len( solution) - 1]
        for j := 0; j < len( lastRow) - 1; j++ {
            
            newRow = append( newRow, lastRow[ j] + lastRow[ j + 1])
        }
        newRow = append( newRow, 1)
        solution = append( solution, newRow)
    }
    return solution
}