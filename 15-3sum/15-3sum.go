func threeSum(nums []int) [][]int {
    
    sort.Ints( nums)
    solutions := make( [][]int, 0)
    for i := 0; i < len(nums) - 2; i++ {
        if i > 0 && nums[ i] == nums[ i - 1] {
            continue
        }
        
        l := i + 1
        r := len(nums) - 1
        for l < r {
            sum := nums[ i] + nums[ l] + nums[ r]
            if sum < 0 {
                
                l++
                
            } else if sum > 0 {
                
                r--
                
            } else {
                
                solutions = append( solutions, []int{ nums[ i], nums[ l], nums[ r]})
                for l < r && nums[ l] == nums[ l + 1] {
                    l++
                }
                for l < r && nums[ r] == nums[ r - 1] {
                    r--
                }
                l++
                r--
            }
        }
    }
    return solutions
}