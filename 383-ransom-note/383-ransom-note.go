func canConstruct(ransomNote string, magazine string) bool {
 
    ransomLetters := make( map[rune]int)
    magLetters := make( map[rune]int)
    
    countLetters( ransomNote, ransomLetters)
    countLetters( magazine, magLetters)
    
    for key, count := range ransomLetters {
        
        if magCount, ok := magLetters[ key]; !ok || magCount < count {
            return false
        }
    }
    return true
}

func countLetters( str string, strLetters map[rune]int) {
    
    for _, char := range str {
        
        strLetters[char]++
    }
}