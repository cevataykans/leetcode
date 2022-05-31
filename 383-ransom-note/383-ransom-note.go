func canConstruct(ransomNote string, magazine string) bool {
 
    charMap := make( map[rune]int )
    for _, ch := range magazine {
        
        charMap[ ch]++
    }
    
    for _, ch := range ransomNote {
        
        charMap[ ch]--
        if charMap[ ch] < 0 {
            
            return false
        }
    }
    return true
}