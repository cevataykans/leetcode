func maxArea(height []int) int {
    
    l := 0
    r := len(height) - 1
    sum := 0
    for l < r {
        
        curSum := (r - l) * min( height, l, r)
        if curSum > sum {
            sum = curSum
        }
        
        if height[ l] < height[ r] {
            l++
        } else {
            r--
        }
    }
    return sum
}

func min( nums []int, i, j int) int {
    if nums[ i] < nums[ j] {
        return nums[ i]
    }
    return nums[ j]
}