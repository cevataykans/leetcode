func checkInclusion(s1 string, s2 string) bool {
    
    if len( s1) > len( s2) {
        
        return false
    }
    
    charFreqs := [26]byte{}
    for _, char := range s1 {
        
        charFreqs[ char - 'a']++
    }
    
    start, end := 0, len(s1) - 1
    curFreqs := [26]byte{}
    for i := 0; i < end; i++ {
        
        curFreqs[ s2[ i] - 'a']++
    }
    
    for end < len( s2) {
        
        curFreqs[ s2[ end] - 'a']++
        if isAnagram( charFreqs, curFreqs) {
            
            return true
        }
        
        curFreqs[ s2[ start] - 'a']--
        start++
        end++
    }
    return false
}

func isAnagram( a1, a2 [26]byte) bool {
    
    for i := range a1 {
        
        if a1[ i] != a2[ i] {
            
            return false
        }
    }
    return true
}