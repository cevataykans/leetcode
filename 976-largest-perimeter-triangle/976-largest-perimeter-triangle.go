func largestPerimeter(nums []int) int {
    
    sort.Ints( nums)
    for i := len( nums) - 1; i >= 2; i-- {
        
        a, b, c := nums[ i], nums[ i - 1], nums[ i - 2]
        if a + b > c && int( math.Abs( float64(a - b) ) ) < c {
            
            return a + b + c
        }
    }
    return 0
}