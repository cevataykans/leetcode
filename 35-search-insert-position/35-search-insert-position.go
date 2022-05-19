func searchInsert( nums []int, target int) int {
    
    return searchInsertPos( nums, target, 0, len( nums) - 1)
}

func searchInsertPos( nums []int, target int, start int, end int ) int {
    
    if start >= end {
        if nums[ start] >= target {
            return start
        }
        return start + 1
    }
    
    middle := ( start + end ) / 2
    if nums[ middle] == target {
        return middle
    }
    
    if nums[ middle] < target {
        return searchInsertPos( nums, target, middle + 1, end)
    }
    return searchInsertPos( nums, target, start, middle - 1)
}