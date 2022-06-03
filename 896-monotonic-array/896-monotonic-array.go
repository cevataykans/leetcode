func isMonotonic(nums []int) bool {
    
    if len( nums) <= 2 {
        
        return true
    }
    
    isGreat, isLow := false, false
    for i := 1; i < len( nums); i++ {
        
        if nums[ i] > nums[ i - 1] {
            
            isGreat = true
        } else if nums[ i] < nums[ i - 1] {
            
            isLow = true
        }
        
        if isGreat && isLow {
            
            return false
        }
    }
    return true
}