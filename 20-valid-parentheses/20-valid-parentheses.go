func isValid(s string) bool {
    
    stack := []rune{}
    
    for _, char := range s {
        
        if char == '}' || char == ']' || char == ')' {
            
            if len( stack) == 0 {
                
                return false
            }
            
            lastChar := stack[ len( stack) - 1]
            if char == ']' && lastChar != '[' || char == '}' && lastChar != '{' || char == ')' && lastChar != '(' {
                
                return false
            } 
            
            stack = stack[ :len( stack) - 1]
            
        } else {
            
            stack = append( stack, char)
        }
    }
    return len( stack) == 0
}