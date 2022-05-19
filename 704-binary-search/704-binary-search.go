func search(nums []int, target int) int {
    
    return searchBinary( nums, target, 0, len(nums) - 1)
}

func searchBinary( nums []int, target int, start int, end int) int {
    
    if start > end {
        return -1
    }
    
    middle := (start + end) / 2
    if nums[ middle] == target {
        return middle
    }
    
    if nums[ middle] < target {
        return searchBinary( nums, target, middle + 1, end)
    }
    return searchBinary( nums, target, start, middle - 1)
}