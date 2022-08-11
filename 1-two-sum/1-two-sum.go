func twoSum(nums []int, target int) []int {
    
    pairs := make( map[int]int)
    for i, num := range nums {
        
        if index, ok := pairs[ target - num]; ok {
            return []int{ i, index }
        }
        pairs[ num] = i
    }
    return nil
}