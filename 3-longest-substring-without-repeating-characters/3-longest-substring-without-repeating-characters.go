func lengthOfLongestSubstring(s string) int {
    
    head := 0
    tail := 0
    maxLength := 0
    counts := make(map[rune]int)
    for head < len(s) {
        
        curRune := rune(s[head])
        if counts[ curRune] > 0 {
            
            tailRune := rune(s[tail])
            counts[ tailRune]--
            tail++
            
        } else {
            
            counts[ curRune]++
            head++
        }
        
        if head - tail > maxLength {
            maxLength = head - tail
        }
    }
    return maxLength
}