func search(nums []int, target int) int {
    
    left, right := 0, len( nums) - 1
    for left <= right {
        
        middle := left + (right - left) / 2
        if nums[ middle] == target {
            
            return middle
        }
        
        var comparator int
        if target < nums[ 0] && nums[ middle] < nums[ 0] || target >= nums[ 0] && nums[ middle] >= nums[ 0] {
            
            comparator = nums[ middle]
            
        } else {
            
            if target < nums[ 0] {
                
                comparator = math.MinInt
            } else {
                
                comparator = math.MaxInt
            }
        }
        
        if target == comparator {
            
            return middle
        } else if target > comparator {
            
            left = middle + 1
        } else {
            
            right = middle - 1
        }
        
    }
    return -1
}