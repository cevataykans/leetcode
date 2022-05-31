func isAlienSorted(words []string, order string) bool {
    
    orderMap := [26]int{}
    curI := 0
    for i := range order {
        
        orderMap[ order[ i] - 'a'] = curI
        curI++
    }
    
    for i := 1; i < len( words); i++ {
        
        if !lexical( words[ i - 1], words[ i], orderMap) {
            
            return false
        }
    }
    return true
}

func lexical( first string, sec string, order [26]int) bool {
    
    i := 0
    j := 0
    
    for i < len( first) && j < len( sec) {
        
        if order[ first[ i] - 'a'] < order[ sec[ i] - 'a'] {
            
            return true
        } 
        
        if order[ first[ i] - 'a'] > order[ sec[ i] - 'a'] {
            
            return false
        }
        
        i++
        j++
    }
    return len( first) <= len( sec)
}