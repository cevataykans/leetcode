func trap(height []int) int {
    
    left := 0
    right := len(height) - 1
    leftMax := -1
    rightMax := -1
    rain := 0
    
    for left < right {
        
        if height[ left] < height[ right] {
            
            if height[ left] > leftMax {
                leftMax = height[ left]
            } else {
                rain += leftMax - height[ left]
            }
            left++
               
        } else {
            
            if height[ right] > rightMax {
                rightMax = height[ right]
            } else {
                rain += rightMax - height[ right]
            }
            right--
        }
    }
    return rain
}