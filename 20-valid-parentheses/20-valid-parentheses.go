func isValid(s string) bool {
    
    stack := []rune{}
    match := map[rune]rune{
        '}' : '{',
        ']' : '[',
        ')' : '(',
    }
    
    for _, char := range s {
        
        val, ok := match[ char]
        if ok {
            
            if len( stack) == 0 {
                
                return false
            }
            
            lastChar := stack[ len( stack) - 1]
            if val != lastChar {
                
                return false
            } 
            
            stack = stack[ :len( stack) - 1]
            
        } else {
            
            stack = append( stack, char)
        }
    }
    return len( stack) == 0
}