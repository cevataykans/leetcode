func lengthOfLongestSubstring(s string) int {
    
    if len(s) == 0 || len( s) == 1 {
        
        return len( s)
    }
    
    charSet := make( map[byte]bool)
    startI, endI := 0, 0
    maxLength := -1
    
    for endI < len( s) {
        
        if charSet[ s[ endI]] {
            
            for s[ startI] != s[ endI] {
                
                delete( charSet, s[ startI])
                startI++
            }
            
            startI++
            
        } else {
            
            charSet[ s[ endI]] = true
        }
        
        if len( charSet) > maxLength {
            
            maxLength = len( charSet)
        }
        endI++
    }
    return maxLength
}