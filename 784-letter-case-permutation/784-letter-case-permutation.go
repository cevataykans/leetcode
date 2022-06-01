func letterCasePermutation(s string) []string {
    
    solution := []string{}
    rawStr := []rune( s)
    solution = findPermutations( rawStr, 0, solution)
    return solution
}

func findPermutations( s []rune, i int, solution []string) []string {
    
    if i == len( s) {
        
        solution = append( solution, string( s))
        return solution
    }
    
    curChar := s[ i]
    if unicode.IsLower( curChar) {
        
        s[ i] = unicode.ToUpper( s[ i])
    } else {
        
        s[ i] = unicode.ToLower( s[ i])
    }
    solution = findPermutations( s, i + 1, solution)
    if unicode.IsLetter( s[ i]) {
        
        s[ i] = curChar
        solution = findPermutations( s, i + 1, solution)
    } 
    return solution
}