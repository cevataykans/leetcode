func twoSum(nums []int, target int) []int {
    
    prevNums := make( map[ int]int)
    for i, value := range nums {
        
        remainingSum := target - value
        index, ok := prevNums[ remainingSum]
        if ok {
            
            return []int{ i, index}
        }
        prevNums[ value] = i
    }
    return nil
}