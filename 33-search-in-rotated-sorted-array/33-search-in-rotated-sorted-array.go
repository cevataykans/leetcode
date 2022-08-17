func search(nums []int, target int) int {
    
    return searchReverseBinary( nums, 0, len(nums) - 1, target)
}

func searchReverseBinary( nums []int, l, r, target int) int {
    
    if l > r {
        return -1
    }
    
    mid := (r - l) / 2 + l
    var comp int
    if (target >= nums[ 0]) == (nums[mid] >= nums[ 0]) {
        
        comp = nums[ mid]
        
    } else {
        
        if nums[ 0] > target {
            comp = math.MinInt
        } else {
            comp = math.MaxInt
        }
    }
    
    if comp == target {
        return mid
    } 
    
    if comp < target {
        return searchReverseBinary( nums, mid + 1, r, target)
    }
    return searchReverseBinary( nums, l, mid - 1, target)
}