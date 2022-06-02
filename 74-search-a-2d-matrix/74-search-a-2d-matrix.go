func searchMatrix(matrix [][]int, target int) bool {
    
    for i := 0; i < len( matrix); i++ {
        
        if matrix[ i][ len( matrix[ i]) - 1] >= target {
            
            return binarySearch( matrix[ i], 0, len( matrix[ i]) - 1, target)
        }
    }
    return false
}

func binarySearch( arr []int, start int, end int, target int) bool {
    
    if start > end {
        
        return false
    }
    
    middle := start + (end - start) / 2
    if arr[ middle] == target {
        
        return true
    }
    
    if arr[ middle] > target {
        
        return binarySearch( arr, start, middle - 1, target)
    }
    return binarySearch( arr, middle + 1, end, target)
}